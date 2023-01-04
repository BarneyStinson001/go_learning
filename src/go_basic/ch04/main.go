package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person struct {
	ID int `uri:"id" binding:"required"`//string类型，后面用,uuid类型，可以返回404  。，真正需求 mongodb mysql 字符串或者int
	Name string `uri:"name" binding:"required"`
}

func main()	  {
	router := gin.Default()
	router.GET("/:name/:id", func(c *gin.Context) {
		var person Person
		if err:=c.ShouldBindUri(&person);err!=nil{
			c.Status(404)
			return//如果没有return 会继续执行下面的语句，返回成功
		}
		c.JSON(http.StatusOK,gin.H{
			"name":person.Name,
			"id":person.ID,
		})


	})
	router.Run(":8084")
}