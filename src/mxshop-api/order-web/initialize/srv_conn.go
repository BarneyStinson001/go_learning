package initialize

import (
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver" // It's important//一定要引入
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"mxshop-api/order-web/global"
	"mxshop-api/order-web/proto"
)


func InitSrvConn(){
	consulInfo := global.ServerConfig.ConsulInfo
	zap.S().Debug(consulInfo,global.ServerConfig.OrderSrvInfo,global.ServerConfig.GoodsSrvConfig)
	goodsconn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s",consulInfo.Host,consulInfo.Port, global.ServerConfig.GoodsSrvConfig.Name),
		//"consul://192.168.18.160:8500/goods-srv?wait=14s&python=srv",
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`), //关于配置看grpc官方文档,
	)
	if err != nil {
		zap.S().Fatal(err)
		zap.S().Fatal("[InitSrvConn] 连接 【商品服务失败】")
	}
	//defer conn.Close()
	goodsSrvClient := proto.NewGoodsClient(goodsconn)
	global.GoodsSrvClient = goodsSrvClient

	orderconn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s",consulInfo.Host,consulInfo.Port, global.ServerConfig.OrderSrvInfo.Name),
		//"consul://192.168.18.160:8500/goods-srv?wait=14s&python=srv",
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`), //关于配置看grpc官方文档,
	)
	if err != nil {
		//log.Fatal(err)
		zap.S().Fatal("[InitSrvConn] 连接 【订单服务失败】")
	}
	//defer conn.Close()
	orderSrvClient := proto.NewOrderClient(orderconn)
	global.OrderSrvClient = orderSrvClient

}




