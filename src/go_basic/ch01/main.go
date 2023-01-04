package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func pong (c *gin.Context){
	c.JSON(http.StatusOK,gin.H{"message":"ping",})
}

func main() {
	//实例化一个对象  flask也差不多 ，还有装饰器，语法更简单
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {//匿名函数
		c.JSON(http.StatusOK, gin.H{//H等于map[string]string  或者map[string]interface{}
		"message": "pong",
		})
	})
	r.GET("/pong",pong)

	r.Run(":8083") // listen and serve on 0.0.0.0:8080

}