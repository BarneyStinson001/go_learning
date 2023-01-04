package initialize

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mxshop-api/user-web/middlewares"
	router2 "mxshop-api/user-web/router"
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
	ApiGroup := router.Group("/u/v1")
	router2.InitUserRouter(ApiGroup)
	router2.InitBaseRouter(ApiGroup)


	return router
}