package initialize

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mxshop-api/order-web/middlewares"
	router2 "mxshop-api/order-web/router"
	"net/http"
)

func Routers() (*gin.Engine) {
	router := gin.Default()

	//监控检查
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"code":http.StatusOK,
			"success":true,
		})
	})
	//配置跨域
	router.Use(middlewares.Cors())
	zap.S().Infof("配置相关url")
	ApiGroup := router.Group("/o/v1")
	router2.InitShopCartRouter(ApiGroup)
	router2.InitOrderRouter(ApiGroup)


	return router
}