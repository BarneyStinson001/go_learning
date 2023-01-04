package router

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mxshop-api/user-web/api"
	"mxshop-api/user-web/middlewares"
)


//main里简单点
func InitUserRouter(group *gin.RouterGroup)  {//也是需要传递指针类型
	//传入Router
	//router :=gin.Default() //不然在每个router里都实例化一个
	UserRouter :=group.Group("user")
	{
		zap.S().Info("配置用户相关接口")
		UserRouter.GET("list", middlewares.JWTAuth(),middlewares.IsAdminAuth(),api.GetUserList)
		UserRouter.POST("passwordlogin",api.PasswordLogin)
		UserRouter.POST("sendsms",api.SendSms)//get请求不行，其实不是和这个相关，当时validatie的错误有handleGRPCerror
		UserRouter.POST("register",api.Register)
	}
}
