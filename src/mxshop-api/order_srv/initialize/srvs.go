package initialize

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"mxshop-api/order_srv/global"
	"mxshop-api/order_srv/proto"
)

func InitSrvs() {
	//初始化第三方微服务的client
	c := global.ServerConfig.ConsulInfo
	goodConn,err:=grpc.Dial(fmt.Sprintf("consul://%s:%d/%s?wait=14s",c.Host,c.Port,global.ServerConfig.GoodsSrvInfo.Name),
		grpc.WithInsecure(),grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
	)
	if err!=nil {
		zap.S().Fatal("[InitSrvsConn]连接 【商品服务失败】")
	}
	global.GoodsSrvClient=proto.NewGoodsClient(goodConn)

	//inventory := global.ServerConfig.ConsulInfo
	inventoryConn,err:=grpc.Dial(fmt.Sprintf("consul://%s:%d/%s?wait=14s",c.Host,c.Port,global.ServerConfig.InventorySrvInfo.Name),
		grpc.WithInsecure(),grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
	)
	if err!=nil{
		zap.S().Fatal("[InitSrvsConn]连接 【库存服务失败】")
	}
	global.InventorySrvClient=proto.NewInventoryClient(inventoryConn)

	global.GoodsSrvClient=proto.NewGoodsClient(goodConn)
}
