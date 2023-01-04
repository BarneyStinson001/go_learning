package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"math/rand"
	"mxshop-api/user-web/forms"
	"mxshop-api/user-web/global"
	"net/http"
	"strings"
	"time"
)

func GenerateSmsCode(width int) string {
	//生成返回width宽度的验证码
	//width从config读取
	numbers :=[10]byte{0,1,2,3,4,5,6,7,8,9}
	r :=len(numbers)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i:=0; i<width;i++{
		fmt.Fprintf(&sb,"%d",numbers[rand.Intn(r)])
	}
	return sb.String()
}

func SendSms(ctx *gin.Context) {
	//发送验证码处理逻辑:
	//1、对表单SendSmsForm进行校验:在config加struct 在global里读
	zap.S().Info("start=============")
	sendSmsForm:= forms.SendSmsForm{}
	if err:=ctx.ShouldBind(&sendSmsForm);err!=nil{
		//HandlerGrpcErrorToHttp(err,ctx)//用下面的函数转换才是正确的,不然表单验证错误的情况下返回200
		HandlerValidatorError(ctx,err)
		//失败要return
		//zap.S().Info(err.Error())
		return
	}
	width:=global.ServerConfig.SMSInfo.Width
	code :=GenerateSmsCode(width)
	//模拟发送短信
	zap.S().Infof("向 %s 用户发送验证码 %s",sendSmsForm.Mobile,code)

	//存入redis,过期时间可配
	duration:=global.ServerConfig.RedisInfo.Expire
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",global.ServerConfig.RedisInfo.Host,global.ServerConfig.RedisInfo.Port),
	})

	rdb.Set(sendSmsForm.Mobile,code,time.Duration(duration)*time.Second)

	ctx.JSON(http.StatusOK,gin.H{
		"msg":"发送成功"+code,
	})
	//time.Sleep(15*time.Second)
	readvalue:=rdb.Get(sendSmsForm.Mobile)//可以用下面的Result方法
	//readvalue:=rdb.Get(sendSmsForm.Mobile).Result()
	//if readvalue == redis.Nil {
		zap.S().Infof("读取数据结果为：%s",readvalue)//超时读取为nil
	//}
}