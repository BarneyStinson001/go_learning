package router

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"mxshop-api/order-web/api/shop_cart"
	"mxshop-api/order-web/middlewares"
)


//main里简单点
func InitShopCartRouter(group *gin.RouterGroup)  {//也是需要传递指针类型
	//传入Router
	//router :=gin.Default() //不然在每个router里都实例化一个
	ShopCartRouter :=group.Group("shopcarts").Use(middlewares.JWTAuth())
	{
		//zap.S().Info("配置用户相关接口")
		ShopCartRouter.GET("", shop_cart.List)
		ShopCartRouter.DELETE("/:id", shop_cart.Delete)
		ShopCartRouter.POST("/:id", shop_cart.New)
		ShopCartRouter.PATCH("/:id", shop_cart.Update)

	}
}
