package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"net/http"
)

var store = base64Captcha.DefaultMemStore

func GetCaptcha(c *gin.Context)  {
	driver:= base64Captcha.NewDriverDigit(80,240,5,0.7,80)
	cp:=base64Captcha.NewCaptcha(driver,store)
	id,base64info,err:=cp.Generate()
	if err!=nil{
		zap.S().Errorf("生成验证码错误：%s",err.Error())//内部打印下错误
		c.JSON(http.StatusInternalServerError,gin.H{
			"msg":"生成验证码错误",//外部不用显示具体的错误
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"captchaId":id,
		"captchaInfo":base64info,
	})
}
