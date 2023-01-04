package global

import (
	ut "github.com/go-playground/universal-translator"
	"mxshop-api/order-web/config"
	"mxshop-api/order-web/proto"
)

var (
	ServerConfig *config.ServerConfig = &config.ServerConfig{}//这边初始化完成，那边不用初始化
	NacosConfig config.NacosConfig

	Trans ut.Translator//a全局变量，包内要导出去需要大写

	GoodsSrvClient proto.GoodsClient
	OrderSrvClient proto.OrderClient

)