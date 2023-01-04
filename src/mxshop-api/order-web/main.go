package main

import (
	"fmt"
	_ "github.com/gin-gonic/gin/binding"
	_ "github.com/hashicorp/consul/api"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"mxshop-api/order-web/global"
	"mxshop-api/order-web/initialize"
	"mxshop-api/order-web/utils"
	consul "mxshop-api/order-web/utils/register"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//port:=8021//这个也放在配置文件里
	//日志也可以放在初始化里面
	//logger,_:=zap.NewDevelopment()
	//zap.ReplaceGlobals(logger)
	initialize.InitLogger()
	initialize.InitConfig()
	port := global.ServerConfig.Port
	zap.S().Infof("打印port%d", port) //注意占位符使用
	//	1、初始化routers
	routers := initialize.Routers()
	zap.S().Info("初始化router ")


	//服务发现：初始化和srv的连接
	initialize.InitSrvConn()

	viper.AutomaticEnv()
	debug := initialize.GetEnvInfo("msshop_debug")
	if !debug {
		port, err := utils.GetFreePort()
		if err == nil {
			global.ServerConfig.Port = port
		}
	}

	//服务注册
	//register_client:=consul.New

	//日志启动
	//logger,_:=zap.NewProduction()
	//defer logger.Sync()
	//sugar := logger.Sugar()
	//上面等价于下面
	//zap.S().Infof("启动服务器，端口：%d",port)//不打印
	//S()可以获取一个全局安全的sugar,可以让我们设置。以及zap.L()

	zap.S().Infof("info信息  启动服务器，端口：%d", port)
	zap.S().Debugf("debug信息 启动服务器，端口：%d", port)


	registerClient:=consul.NewRegisteryClient(global.ServerConfig.ConsulInfo.Host,global.ServerConfig.ConsulInfo.Port)//传入consul的配置
	serviceId:=fmt.Sprintf("%s",uuid.NewV4())
	err:=registerClient.Register(global.ServerConfig.Host,global.ServerConfig.Port,global.ServerConfig.Name,global.ServerConfig.Tags,serviceId)
	if err!=nil{
		zap.S().Panic("注册失败", err.Error())
	}

	go func(){
		if err := routers.Run(fmt.Sprintf(":%d", port)); err != nil {
			zap.S().Panic("启动失败", err.Error())
		}
	}()

	//接收终止信号。退出后注销服务
	quit:=make(chan  os.Signal)
	signal.Notify(quit,syscall.SIGINT,syscall.SIGTERM)
	<-quit
	zap.S().Info("注销")

	if err = registerClient.Deregister(serviceId);err!=nil{
		zap.S().Info("注销失败:", err.Error())
	}else{
		zap.S().Info("注销成功")
	}
}
