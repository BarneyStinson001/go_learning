package router

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mxshop-api/goods-web/api"
)


//main里简单点
func InitGoodsRouter(group *gin.RouterGroup)  {//也是需要传递指针类型
	//传入Router
	//router :=gin.Default() //不然在每个router里都实例化一个
	GoodsRouter :=group.Group("goods")
	{
		zap.S().Info("配置用户相关接口")
		GoodsRouter.GET("list", api.GetGoodsList)
	}
}
