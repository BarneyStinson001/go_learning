package ch04

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	goodsGroup := router.Group("/goods")
	{
		goodsGroup.GET("/",goodsList)
		goodsGroup.GET("/:id/:action",goodsDetail)//上面用list，会冲突。  list当成id传入
		goodsGroup.POST("/add",createGoods)//add也可以不要，直接用post方法
	}
	router.Run(":8083")
}

func createGoods(context *gin.Context) {
	
}

func goodsDetail(context *gin.Context) {
	id := context.Param("id")
	action:=context.Param("action")
	context.JSON(http.StatusOK,gin.H{
		"id":id,
		"action":action,
	})
}

func goodsList(context *gin.Context) {
	context.JSON(http.StatusOK,gin.H{
		"list":"a",
	})
}