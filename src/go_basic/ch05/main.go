package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//获取get参数和post参数
//http://localhost:8083/welcome?
//http://localhost:8083/welcome?firstname=ella&lastname=frank
//post参数获取
//requests.post("http://127.1:8083/form_post",data = {"message“：”1234342。“id":'alice'})
//post?id=123&page=12    data={    "message": "1028endcivd4", "name": "jack",}



func main() {
	router :=gin.Default()
	router.GET("/welcome",welcome)
	router.POST("form_post",form_post)

	router.POST("/post",getpost)
	router.Run(":8083")
}
//既有get，又有post
func getpost(context *gin.Context) {
	id := context.Query("id")
	page := context.DefaultQuery("page","0")
	name := context.PostForm("name")
	message := context.DefaultPostForm("message","信息")

	context.JSON(http.StatusOK,gin.H{
		"id":id,
		"page":page,
		"name":name,
		"message":message,
	})
}

func form_post(context *gin.Context) {
	message := context.PostForm("message")
	nick :=context.DefaultPostForm("nick","anony")

	context.JSON(http.StatusOK,gin.H{
		"message":message,
		"nick":nick,
	})
}

func welcome(context *gin.Context) {
	firstName := context.DefaultQuery("firstname","alice")
	lastName := context.DefaultQuery("lastname","bob")
	context.JSON(http.StatusOK,gin.H{
		"firstname":firstName,
		"lastname":lastName,
	})
}
