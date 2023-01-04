package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK,gin.H{
			"message":"pong",
		})
	})
	go func() {
		router.Run(":8003")//夯住，阻塞
	}()

	//如果想要接收到信号， chanmal
	quit :=make(chan os.Signal)
	signal.Notify(quit,syscall.SIGINT,syscall.SIGTERM)
	<-quit

	//
	fmt.Println("关闭server中")
	fmt.Println("注销服务")
}