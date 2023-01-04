package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)


//自定义中间件
func MyLogger()  gin.HandlerFunc {
	return func(context *gin.Context) {
		t :=time.Now()
		//前处理
		context.Set("example","123456")
		//原来的逻辑继续执行
		context.Next()
		//后处理
		end:=time.Since(t)
		fmt.Println("耗时：%V\n",end)
		status:=context.Writer.Status()
		fmt.Println("状态",status)
	}
}
func Tokenauthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		for k,v :=range c.Request.Header{
			fmt.Println(k,v)
			if k=="X-Token" {
				token=v[0]
			}
		}
		//time.Sleep(time.Second)
		print("token is: ",token)
		if token != "alice"{
			c.JSON(http.StatusUnauthorized,gin.H{
				"message":"you are not allowed",
			})
			//return//不起作用
			c.Abort()
		}
		c.Next()
	}
}



//http://127.0.0.1:8009/ping
func main() {
	//router  := gin.Default()
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	//	以上等于Default  全局作用
	//想要仅作用在某些接口
	authorized :=router.Group("/goods")
	authorized.Use(AuthRequired)

	router.Use(MyLogger())
	router.Use(Tokenauthentication())
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK,gin.H{
			"message":"pong",
		})
	})

	router.Run(":8009")

}

func AuthRequired(context *gin.Context) {
	return
}