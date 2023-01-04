package global

import (
	ut "github.com/go-playground/universal-translator"
	"mxshop-api/goods-web/config"
	"mxshop-api/goods-web/proto"
)

var (
	ServerConfig *config.ServerConfig = &config.ServerConfig{}//这边初始化完成，那边不用初始化
	Trans ut.Translator//a全局变量，包内要导出去需要大写

	GoodsSrvClient proto.GoodsClient

)