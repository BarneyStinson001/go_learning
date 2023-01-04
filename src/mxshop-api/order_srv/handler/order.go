package handler

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"math/rand"
	"mxshop-api/order_srv/global"
	"mxshop-api/order_srv/model"
	"mxshop-api/order_srv/proto"
	"time"
)

//一个个做
type OrdersServer struct {
	proto.UnimplementedOrderServer
}

func GenerateOrderSn(userId int32) string {
	//	年月日时分秒  userId  随机数
	now := time.Now()
	rand.Seed(time.Now().UnixNano())
	orderSn := fmt.Sprintf("%d%d%d%d%d%d%d%d",
		now.Year(), now.Month(), now.Day, now.Hour(), now.Minute(), now.Nanosecond(),
		userId, rand.Intn(90)+10)
	return orderSn
}

func (*OrdersServer) CartItemList(ctx context.Context, req *proto.UserInfo) (*proto.CartItemListRsp, error) {
	//获取当前用户的购物车列表
	var shopCarts []model.ShoppingCart

	res := global.DB.Where(&model.ShoppingCart{User: req.Id}).Find(&shopCarts)
	if res.Error != nil {
		return nil, res.Error
	}
	rsp := proto.CartItemListRsp{
		Total: int32(res.RowsAffected),
	}
	for _, shopCart := range shopCarts {
		rsp.Data = append(rsp.Data, &proto.ShopCartInfoRsp{
			Id:      shopCart.ID,
			UserId:  shopCart.User,
			GoodsId: shopCart.Goods,
			Nums:    shopCart.Nums,
			Checked: shopCart.Checked,
		})
	}
	return &rsp, nil
}

func (*OrdersServer) CreateCartItem(ctx context.Context, req *proto.CartItemReq) (*proto.ShopCartInfoRsp, error) {
	//加入购物车，从0到1，已有商品增加需要合并
	var shopCart model.ShoppingCart

	res := global.DB.Where(&model.ShoppingCart{
		User:  req.UserId,
		Goods: req.GoodsId,
	}).First(&shopCart)

	if res.RowsAffected == 1 {
		//已存在;合并记录，更新操作
		shopCart.Nums += req.Nums
	} else {
		//插入操作
		shopCart.User = req.UserId
		shopCart.Goods = req.GoodsId
		shopCart.Nums = req.Nums
		shopCart.Checked = false //默认不选中
	}
	global.DB.Save(&shopCart)
	return &proto.ShopCartInfoRsp{Id: shopCart.ID}, nil

}

func (*OrdersServer) UpdateCartItem(ctx context.Context, req *proto.CartItemReq) (*emptypb.Empty, error) {
	//更新：数量，勾选状态
	//先查询
	var shopCart model.ShoppingCart
	//res:=global.DB.Where(&model.ShoppingCart{User: req.UserId,Goods: req.GoodsId,}).First(&shopCart)
	res := global.DB.First(&shopCart, req.Id) //为啥不加用户Id查询

	if res.RowsAffected == 0 {
		//更新可以参考前面的goods,库存等
		return nil, status.Errorf(codes.NotFound, "商品不在购物车")
	}
	if req.Nums > 0 {
		shopCart.Nums = req.Nums
	}
	shopCart.Checked = req.Checked //小心protobuf传0
	global.DB.Save(&shopCart)
	return &emptypb.Empty{}, nil

}

func (*OrdersServer) DeleteCartItem(ctx context.Context, req *proto.CartItemReq) (*emptypb.Empty, error) {
	if res := global.DB.Delete(&model.ShoppingCart{}, req.Id); res.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品在购物车不存在")
	}
	return &emptypb.Empty{}, nil
}

func (*OrdersServer) OrderList(ctx context.Context, req *proto.OrderFilterReq) (*proto.OrderListRsp, error) {

	var orders []model.OrderInfo
	var rsp proto.OrderListRsp
	//后台管理系统 ，电商系统中心，底层不关心。业务传不传用户Id
	//管理员查询 不传userId, 默认值0,在gorm里会被忽略
	//res:=global.DB.Where(&model.OrderInfo{User: req.UserId}).Find(&orders)//为了total，查询所有存到orders
	var total int64
	global.DB.Where(&model.OrderInfo{User: req.UserId}).Count(&total) //为了total，查询所有存到orders
	rsp.Total = int32(total)

	//分页，也要带上用户过滤
	global.DB.Scopes(Paginate(int(req.Pages), int(req.NumsPerPage))).Find(&orders)
	for _, order := range orders {
		rsp.Data = append(rsp.Data, &proto.OrderInfoRsp{
			Id:      order.ID,
			UserId:  order.User,
			OrderSn: order.OrderSn,
			PayType: order.PayType,
			Status:  order.Status,
			Post:    order.Post,
			Total:   order.OrderMount,
			Address: order.Address,
			Name:    order.SignerName,
			Mobile:  order.SingerMobile,
			AddTime: "",
		})
	}
	return &rsp, nil
	//不带商品详情，不然需要join
}

func (*OrdersServer) OrderDetail(ctx context.Context, req *proto.OrderReq) (*proto.OrderDetailRsp, error) {
	var order model.OrderInfo
	var rsp proto.OrderDetailRsp
	//需要注意权限，这个订单是否为当前用户的。web先查询一下？个人中心是可以的传ID，后台管理系统不行只传递orderId
	//权限也是web层管
	//默认值为0，不拼接
	if res := global.DB.Where(&model.OrderInfo{BaseModel: model.BaseModel{ID: req.Id}, User: req.UserId}).First(&order); res.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "订单不存在")
	}
	//嵌套赋值
	//rsp.OrderInfo.Id = order.ID
	//rsp.OrderInfo.UserId = order.User
	//rsp.OrderInfo.OrderSn = order.OrderSn
	//rsp.OrderInfo.PayType = order.PayType
	//rsp.OrderInfo.Status = order.Status
	//rsp.OrderInfo.Post = order.Post
	//rsp.OrderInfo.Total = order.OrderMount
	//报错，rsp.OrderInfo为nil
	orderInfo:=proto.OrderInfoRsp{}
	orderInfo.Id = order.ID
	orderInfo.UserId = order.User
	orderInfo.OrderSn = order.OrderSn
	orderInfo.PayType = order.PayType
	orderInfo.Status = order.Status
	orderInfo.Post = order.Post
	orderInfo.Total = order.OrderMount
	rsp.OrderInfo=&orderInfo

	var orderGoods []model.OrderGoods
	global.DB.Where(&model.OrderGoods{Order: order.ID}).Find(&orderGoods) //也要判断res.Affected
	//冗余字段，好在不用跨服务查。
	//但商品修改，有可能需要通知其他表的冗余字段同步修改。涉及到分布式事务
	for _, orderGood := range orderGoods {
		rsp.Goods = append(rsp.Goods, &proto.OrderItemRsp{
			GoodsId:    orderGood.Goods,
			GoodsName:  orderGood.GoodsName,
			GoodsImage: orderGood.GoodsImage,
			GoodsPrice: orderGood.GoodsPrice,
			Nums:       orderGood.Nums,
		})
	}
	return &rsp, nil
}

//新建订单，最复杂

func (*OrdersServer) CreateOrder(ctx context.Context, req *proto.OrderReq) (*proto.OrderInfoRsp, error) {
	/*
		新建订单：
		0、从购物车选中的购物车
		1、计算总金额， 不能照前端传递的，得到商品服务区查询。跨微服务调用
		2、下单，库存预扣减
		3、 入表  订单的基本信息  订单表，订单商品信息表
		4、购物车删除已购买的记录
	*/
	var goodIds []int32

	var shopCarts []model.ShoppingCart
	goodsNumsMap := make(map[int32]int32)
	if res := global.DB.Where(&model.ShoppingCart{User: req.UserId, Checked: true}).Find(&shopCarts); res.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "没有选中的商品")
	}
	//价格。批量查询
	for _, shopCart := range shopCarts {
		goodIds = append(goodIds, shopCart.Goods)
		goodsNumsMap[shopCart.Goods] = shopCart.Nums
	}

	//跨服务调用
	//批量获取
	goods, err := global.GoodsSrvClient.BatchGetGoods(context.Background(), &proto.BatchGoodsIdInfo{Id: goodIds})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "批量查询商品信息失败")
	}

	//总价  shopCarts里每件商品都有价格   shopCarts遍历不好，goodsNumsMap存储
	var orderGoods []*model.OrderGoods //批量插入
	var orderAmount float32
	var goodsInvInfo []*proto.GoodsInvInfo
	for _, good := range goods.Data {
		orderAmount += good.ShopPrice * float32(goodsNumsMap[good.Id])
		orderGoods = append(orderGoods, &model.OrderGoods{
			Goods:      good.Id,
			GoodsName:  good.Name,
			GoodsPrice: good.ShopPrice,
			Nums:       goodsNumsMap[good.Id],
			GoodsImage: good.GoodsFrontImage,
		})

		goodsInvInfo = append(goodsInvInfo, &proto.GoodsInvInfo{
			GoodsId: good.Id,
			Num:     goodsNumsMap[good.Id],
		})
	}

	//库存扣减，扣减不了不能插订单表
	if _, err = global.InventorySrvClient.Sell(context.Background(), &proto.SellInfo{GoodsInvInfo: goodsInvInfo}); err != nil {
		return nil, status.Errorf(codes.ResourceExhausted, "库存扣减失败")
	}
	//生成订单表 1-13

	tx:=global.DB.Begin()
	order := model.OrderInfo{
		OrderSn:      GenerateOrderSn(req.Id),
		OrderMount:   orderAmount,
		Address:      req.Address,
		SignerName:   req.Name,
		SingerMobile: req.Mobile,
		Post:         req.Post,
		User:			req.UserId,
	}
	//global.DB.Save(&order)//全都换成tx
	if res:=tx.Save(&order);res.RowsAffected==0{
		tx.Rollback()
		return nil, status.Errorf(codes.Internal, "创建订单失败")

	}

	//如果批量插入orderGood表失败
	for _, orderGood := range orderGoods {
		orderGood.Order = order.ID
	}
	if res:=tx.CreateInBatches(orderGoods, 100);res.RowsAffected==0{
		tx.Rollback()
		return nil, status.Errorf(codes.Internal, "创建订单失败")
	}
	//购物车删除
	if res:=tx.Where(&model.ShoppingCart{User: req.UserId,Checked: true}).Delete(model.ShoppingCart{});res.RowsAffected==0{
		res.Rollback()
		return nil, status.Errorf(codes.Internal, "创建订单失败")

	}
	//tx.Commit()

	//本地事务:所有的修改都需要事务。
	//如果rollback,返回失败
	return &proto.OrderInfoRsp{Id:order.ID,OrderSn: order.OrderSn,Total: order.OrderMount},nil


	//分布式事务
}


func (*OrdersServer) UpdateOrderStatus(ctx context.Context,req *proto.OrderStatus) (*emptypb.Empty, error) {
	//支付宝回调后，更新订单状态。
	//需不需要先查询，再更新。  直接update
	//根据orderSn
	if res:=global.DB.Model(&model.OrderInfo{}).Where("order_sn = ?",req.OrderSn).Update("status",req.Status);res.RowsAffected==0{
		return nil, status.Errorf(codes.NotFound, "订单不存在")
	}
	return &emptypb.Empty{},nil
}
