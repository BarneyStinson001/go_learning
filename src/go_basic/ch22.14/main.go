package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

//http://127.0.0.1:8009/index
//http://127.0.0.1:8009/goods
//http://127.0.0.1:8009/users/list
func main() {
	dir,_:=filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Println("path is:",dir)

	router := gin.Default()

	//index.tmpl路径	需要先加载
	router.LoadHTMLFiles("D:\\github_learning\\go_learning\\src\\go_basic\\ch22.14\\template\\index.tmpl","D:\\github_learning\\go_learning\\src\\go_basic\\ch22.14\\template\\goods.html")
	router.LoadHTMLGlob("D:\\github_learning\\go_learning\\src\\go_basic\\ch22.14\\template\\**\\*")
	//多个文件加载用*,见下面相对路径

	//建议使用相当路径，那不就能在goland运行  。golang的临时路径 C:\Users\zhuhangxin\AppData\Local\Temp
	//router.LoadHTMLFiles("template\\index.tmpl")//需要进到文件夹下，运行
	//router.LoadHTMLGlob("template/*")

	router.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK,"index.tmpl",gin.H{
			"title":"Alice's Blog",
		})
	})
	router.GET("/goods", func(context *gin.Context) {
		context.HTML(http.StatusOK,"goods.html",gin.H{
			"name":"Alice's blogs",
		})
	})


	router.GET("/users/list", func(context *gin.Context) {
		//users/list.html不需要路径
		//context.HTML(http.StatusOK,"list.html",gin.H{
		//	"title":"users list",
		//})

		//不同文件夹下的同名文件如何解决同名问题，加别名
		context.HTML(http.StatusOK,"users/list.html",gin.H{
			"title":"users list",
		})
	})
	router.GET("/goods/list", func(context *gin.Context) {
		//users/list.html不需要路径
		//不同文件夹下的同名文件如何解决同名问题，加别名
		context.HTML(http.StatusOK,"goods/list.html",gin.H{
			"title":"goods list",
		})
	})



	router.Run(":8009")
}
