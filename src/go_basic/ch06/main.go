package main

import (
	"github.com/gin-gonic/gin"
	"go_learning/src/go_basic/ch06/proto"
	"net/http"
)

//http://localhost:8083/moreJSON
//{"user":"alice","Message":"test data","Number":20}
func main() {
	router :=gin.Default()

	router.GET("/moreJSON",moreJSON)
	router.GET("/someProtoBuf",returnProto)

	router.Run(":8083")

}

func returnProto(context *gin.Context) {
	courses := []string{"math1","math2",":math3"}
	user := &proto.Teacher{Name: "Dr. alice",Course: courses}
	context.ProtoBuf(http.StatusOK,user)//生成protobuf对象
}
//用浏览器请求会下载原文，http://localhost:8083/someProtoBuf
//可以用python测试，ParseFromString

//json的tag
func moreJSON(context *gin.Context) {
	var msg struct{
		Name string `json:"user"`//json格式中键会转化成user
		Message string
		Number int
	}
	msg.Name="alice"
	msg.Message="test data"
	msg.Number=20

	context.JSON(http.StatusOK,msg)
}