package main

import (
	"google.golang.org/grpc"

	"mxshop-api/order_srv/proto"
)

//var brandClient proto.GoodsClient
var conn *grpc.ClientConn
var orderClient  proto.OrderClient



func Init(){
	var err error
	conn, err = grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	orderClient = proto.NewOrderClient(conn)
}

func Test()  {
	
}

func main() {
	Init()
	//测试订单接口
	//添加商品到购物车
	TestCreateCartItem(userId,nums,goodsId)//检查数据库
	//再次调用,增加数量,返回id不变
	//i查询购物车商品
	TestCartItemList()//循环打印购物车里的商品Id,GoodsId,Nums
	//更新购物车状态  :勾选
	TestUpdateCartItem()//检查数据库
	//新建订单接口
	TestCreateOrder()//测试无选中，测试需要启动商品和库存服务。 检查数据库记录
	//获取订单详情
	TestGetOrderDetail//循环打印订单商品详情
	//订单列表
	TestorderList()//管理员查询和用户查询
	//更新订单状态
	TestUpdateOrder()


	conn.Close()
}

