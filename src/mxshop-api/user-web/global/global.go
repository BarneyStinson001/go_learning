package global

import (
	ut "github.com/go-playground/universal-translator"
	"mxshop-api/user-web/config"
	"mxshop-api/user-web/proto"
)

var (
	ServerConfig *config.ServerConfig = &config.ServerConfig{}//这边初始化完成，那边不用初始化
	Trans ut.Translator//a全局变量，包内要导出去需要大写

	UserSrvClient proto.UserClient

)