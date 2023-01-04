package initialize

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	_ "github.com/mbobakov/grpc-consul-resolver" // It's important//一定要引入
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"mxshop-api/user-web/global"
	"mxshop-api/user-web/proto"
)


func InitSrvConn(){
	//consulInfo := global.ServerConfig.ConsulInfo

	conn, err := grpc.Dial(
		//fmt.Sprintf("consul://%s:%d/%s?wait=14s",consulInfo.Host,consulInfo.Port, global.ServerConfig.UserSrvConfig.Name),
		"consul://192.168.18.160:8500/user-srv?wait=14s&python=srv",
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`), //关于配置看grpc官方文档,
	)
	if err != nil {
		//log.Fatal(err)
		zap.S().Fatal("[InitSrvConn] 连接 【用户服务失败】")
	}
	//defer conn.Close()
	userSrvClient := proto.NewUserClient(conn)
	global.UserSrvClient = userSrvClient
}



//老的
func InitSrvConn2() {
	//初始化grpc服务

	cfg := api.DefaultConfig()
	consulInfo := global.ServerConfig.ConsulInfo
	cfg.Address = fmt.Sprintf("%s:%d", consulInfo.Host, consulInfo.Port)
	//读取拿到consul的ip port

	userSrvHost := ""
	userSrvPort := 0
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	//data, err := client.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%s"`, name))
	data, err := client.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%s"`, global.ServerConfig.UserSrvConfig.Name))//注意配置文件里的符号 -和_
	//data,err:=client.Agent().ServicesWithFilter(`Service == "user-srv"`)

	for _, value := range data {
		userSrvHost = value.Address
		userSrvPort = value.Port
	}
	//以上为服务发现，拿到userSrvHost和userSrvPort

	//如果userSrvHost为空，说明用户服务失败
	if userSrvHost == "" {
		//这里是初始化，没有context，不用返回
		zap.S().Fatal("[InitSrvConn] 连接 【用户服务失败】")
	}

	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", userSrvHost, userSrvPort), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("连接【用户服务】失败", "msg", err.Error())
	}

	userSrvClient := proto.NewUserClient(userConn)
	global.UserSrvClient = userSrvClient
	//还有其他问题：后续用户服务下线了，端口变了，ip变了  ---后面有负载均衡
	//好处，不用再进行建立连接，三次握手。
	//一个连接给多个goroutine共用，会不会影响性能 ---grpc连接池 grpc go pool
}
