package util

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func InitRouter() *gin.Engine {

	r:=gin.Default()


	r.GET("/", func(context *gin.Context) {
		log.Println("received req")
		context.JSONP(http.StatusOK,gin.H{
			"live_dash" : gin.H{
				"runner_num":123,
				"rate":2,
				"playtime":1800,
			},
		})
	})

	r.POST("set",set)
	r.GET("list",list)


	return r
}

