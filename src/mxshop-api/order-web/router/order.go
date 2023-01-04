package router

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"mxshop-api/order-web/api/order"
	"mxshop-api/order-web/middlewares"
)

//main里简单点
func InitOrderRouter(group *gin.RouterGroup) { //也是需要传递指针类型
	//传入Router
	//router :=gin.Default() //不然在每个router里都实例化一个
	OrderRouter := group.Group("orders")
	{
		//zap.S().Info("配置用户相关接口")
		OrderRouter.GET("", middlewares.JWTAuth(),middlewares.IsAdminAuth(),order.List)
		OrderRouter.POST("", order.New)
		OrderRouter.GET("/:id", order.Detail)
	}
}
