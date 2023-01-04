package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"mxshop-api/inventory_srv/global"
	"mxshop-api/inventory_srv/model"
	"mxshop-api/inventory_srv/proto"
	"sync"
)

//一个个做
type  InventoryServer struct{
	proto.UnimplementedInventoryServer
}

func (*InventoryServer) SetInv(ctx context.Context,req *proto.GoodsInvInfo) (*emptypb.Empty, error) {
	//设置库存，更新库存
	var inv model.Inventory
	global.DB.First(&inv,req.GoodsId)
	if inv.Goods == 0 {
		//初次设置
		inv.Goods=req.GoodsId   //其实不管有没有，都可以强制设置一下
	}
	inv.Stocks=req.Num
	global.DB.Save(&inv)

	return &emptypb.Empty{},nil
}


func (*InventoryServer) InvDetail(ctx context.Context,req *proto.GoodsInvInfo) (*proto.GoodsInvInfo, error) {
	var inv model.Inventory
	if  res:=global.DB.First(&inv,req.GoodsId);res.RowsAffected==0{
		return nil,status.Errorf(codes.NotFound,"库存信息不存在")
	}
	return &proto.GoodsInvInfo{
		GoodsId: inv.Goods,
		Num: inv.Stocks,
	},nil
}

var m sync.Mutex
func (*InventoryServer) Sell( ctx context.Context,req *proto.SellInfo) (*emptypb.Empty, error) {
	//事务
	//订单中有不同数量的不同商品
	client := goredislib.NewClient(&goredislib.Options{
		Addr: "192.168.0.104:6379",
	})
	pool := goredis.NewPool(client) // or, pool := redigo.NewPool(...)
	rs := redsync.New(pool)
	//TODO
	tx:=global.DB.Begin()//手动事务
	//m.Lock()
	for _,goods:=range  req.GoodsInvInfo{
		var inv model.Inventory
		//先查询
		//这个First查询用的是主键id，得用where
		//if  res:=global.DB.First(&inv,goods.GoodsId);res.RowsAffected==0{
		//mysql悲观锁
		//if  res:=tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where(&model.Inventory{Goods: goods.GoodsId}).First(&inv);res.RowsAffected==0{
		//	tx.Rollback()//回滚
		//	return nil,status.Errorf(codes.InvalidArgument,"库存信息不存在")
		//}
		//mysql乐观锁
		//for {
		//	if  res:=global.DB.Where(&model.Inventory{Goods: goods.GoodsId}).First(&inv);res.RowsAffected==0{
		//		tx.Rollback()//回滚
		//		return nil,status.Errorf(codes.InvalidArgument,"库存信息不存在")
		//	}
		//	//库存是否充足
		//	if inv.Stocks< goods.Num{
		//		tx.Rollback()//回滚
		//		return nil,status.Errorf(codes.ResourceExhausted,"库存不足")
		//	}
		//	inv.Stocks-=goods.Num
		//	//global.DB.Save(&inv)//改成事务tx
		//	//tx.Save(&inv)
		//	//换成根据id和version更新，update inv set stocks=stocks -2 ,version = version +1 where goods= 421 and version = version
		//	//if res:=tx.Model(&model.Inventory{}).Where("goods = ? and version = ?",goods.GoodsId,inv.Version).Updates(model.Inventory{Version: inv.Version+1,Stocks: inv.Stocks});res.RowsAffected==0{
		//	//零值，不设置//UPDATE `inventory` SET `update_time`='2022-11-05 21:46:42.239',`stocks`=20,`version`=30 WHERE (goods = 505 and version = 29) AND `inventory`.`deleted_at` IS NULL
		//	//强制更新
		//	if res:=tx.Model(&model.Inventory{}).Select("stocks","version").Where("goods = ? and version = ?",goods.GoodsId,inv.Version).Updates(model.Inventory{Version: inv.Version+1,Stocks: inv.Stocks});res.RowsAffected==0{
		//
		//		zap.S().Info("库存扣减失败")
		//	}else{
		//		break
		//	}
		//}
		//redis分布式锁
		//todo
	}
	tx.Commit()//不commit不会保存到数据库
	//m.Unlock()
	return &emptypb.Empty{},nil
}
//改用mysql  for update 悲观锁

func (*InventoryServer) Reback(ctx context.Context,req  *proto.SellInfo) (*emptypb.Empty, error) {
	//事务
	//和扣减反过来，且不需要判断库存够不够减
	tx:=global.DB.Begin()//手动事务

	for _,goods:=range  req.GoodsInvInfo{
		var inv model.Inventory
		//先查询
		if  res:=global.DB.First(&inv,goods.GoodsId);res.RowsAffected==0{
			tx.Rollback()//回滚
			return nil,status.Errorf(codes.InvalidArgument,"库存信息不存在")
		}
		inv.Stocks+=goods.Num
		//global.DB.Save(&inv)//改成事务tx
		tx.Save(&inv)
	}
	tx.Commit()//不commit不会保存到数据库

	return &emptypb.Empty{},nil

}
