package api

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mxshop-api/user-web/forms"
	"mxshop-api/user-web/global"
	"mxshop-api/user-web/global/response"
	"mxshop-api/user-web/middlewares"
	"mxshop-api/user-web/models"
	"mxshop-api/user-web/proto"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//这个函数如果放在initialize的validator里面，就会涉及到循环导入的问题，这边要引入initialize包，显示 import cycle not allowed
func RemoveTopStruct(fields map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fields {
		idx := strings.Index(field, ".")
		rsp[field[idx+1:]] = err
	}
	return rsp
}

func HandlerValidatorError(c *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		c.JSON(http.StatusOK, gin.H{"msg": err.Error()})
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"error": RemoveTopStruct(errs.Translate(global.Trans)),
	})
	return
}

func HandlerGrpcErrorToHttp(err error, c *gin.Context) {
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": e.Message(),
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "内部错误",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "参数错误" + e.Message(),
				})
			case codes.Unavailable:
				c.JSON(http.StatusInternalServerError, gin.H{
					//"msg":"用户服务不可用",
					"msg": e.Message(),
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": e.Code(),
				})
			}
		}
	}
}

func GetUserList(ctx *gin.Context) { //是指针类型
	//发起连接，处理和返回数据
	//验证能不能请求到，打断点，或者是加日志。再用浏览器请求触发
	zap.S().Debug("请求用户列表")

	//到配置文件中去了
	//ip := "127.0.0.1"//写成127.0.0，1，显示no such host
	//port := 50051 //访问python层
	//从配置文件读
	//userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", global.ServerConfig.UserSrvConfig.Host, global.ServerConfig.UserSrvConfig.Port), grpc.WithInsecure())
	//if err != nil {
	//	zap.S().Errorw("[GetUserList] 连接 【用户服务失败】",
	//		"msg", err.Error(),
	//	)
	//}
	//userSrvClient := proto.NewUserClient(userConn)
	//上面都从global里获取
	//print("用户列表请求")
	//25-1，获取参数，没有就给默认值
	PageNo := ctx.DefaultQuery("PageNo", "0")
	pnInt, _ := strconv.Atoi(PageNo)
	PageSize := ctx.DefaultQuery("PageSize", "10")
	pSizeInt, _ := strconv.Atoi(PageSize)

	rsp, err := global.UserSrvClient.GetUserList(context.Background(), &proto.PageInfo{
		PageNo:   uint32(pnInt),
		PageSize: uint32(pSizeInt),
	})
	if err != nil {
		zap.S().Errorw("[GetUserList] 查询【用户列表】 失败")
		HandlerGrpcErrorToHttp(err, ctx)
		return
	}

	//result :=make(map[string]interface{},0)
	result := make([]interface{}, 0)
	for _, val := range rsp.Data { //转化大写了
		//data:= make(map[string]interface{})
		//data["id"] = val.Id
		//data["name"]=val.NickName
		//data["birthday"]=val.Birthday
		//data["gender"]=val.Gender
		//data["mobile"]=val.Mobile
		////不是go的形式

		//user :=response.UserResponse{
		//	Id:val.Id,
		//	NickName: val.NickName,
		//	//Birthday: time.Time(time.Unix(int64(val.Birthday),0)),//2022-07-24T00:00:00+08:00
		//	Birthday: time.Time(time.Unix(int64(val.Birthday),0)).Format("2006-01-02"),
		//	Gender: val.Gender,
		//	Mobile: val.Mobile,
		//}
		//
		user := response.UserResponse{
			Id:       val.Id,
			NickName: val.NickName,
			//Birthday: time.Time(time.Unix(int64(val.Birthday),0)),//2022-07-24T00:00:00+08:00
			Birthday: response.JsonTime(time.Unix(int64(val.Birthday), 0)),
			Gender:   val.Gender,
			Mobile:   val.Mobile,
		}

		result = append(result, user)
	}
	ctx.JSON(http.StatusOK, result)
}

func PasswordLogin(c *gin.Context) {
	//表单验证
	passwordLoginForm := forms.PasswordLoginForm{}
	if err := c.ShouldBind(&passwordLoginForm); err != nil {
		//返回错误信息：翻译，返回的错误格式化
		//errs, ok := err.(validator.ValidationErrors)
		//if !ok {
		//	c.JSON(http.StatusOK, gin.H{"msg": err.Error()})
		//}
		//c.JSON(http.StatusBadRequest, gin.H{
		//	"error": RemoveTopStruct(errs.Translate(global.Trans)),
		//})
		//return
		//把上面的封装成函数使用，需要传递  c err 给HandlerValidatorError
		HandlerValidatorError(c, err)
		return
	}
	//验证码校验
	if !store.Verify(passwordLoginForm.CaptchaId,passwordLoginForm.Captcha,true){
		c.JSON(http.StatusBadRequest,gin.H{
			"captcha":"验证码错误",
		})
		//{"captcha":"验证码错误"}{"expired_at":1661530180000,"id":4,"nick_name":"test003","token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6NCwiTmlja05hbWUiOiJ0ZXN0MDAzIiwiQXV0aG9yaXR5SWQiOjIsImV4cCI6MTY2MTUzMDE4MCwiaXNzIjoiYWxpY2UiLCJuYmYiOjE2NTg5MzgxODB9.lx7y898SS6UR89B74ltRjrcuJPHdsgEvQ35s4QfMuxQ"}
		//需要return
		return
	}


	//跟前面一样查询后台 ，可优化
	//userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", global.ServerConfig.UserSrvConfig.Host,
	//	global.ServerConfig.UserSrvConfig.Port), grpc.WithInsecure())
	//if err != nil {
	//	zap.S().Errorw("[GetUserList]连接 【用户服务】失败", "msg", err.Error())
	//}
	//userSrvClient := proto.NewUserClient(userConn)
	//用户存不存在
	//以上从global获取

	rsp, err := global.UserSrvClient.GetUserByMobile(context.Background(), &proto.MobileRequest{Mobile: passwordLoginForm.Mobile})
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusBadRequest, map[string]string{
					"mobile": "用户不存在",
				})
			default:
				c.JSON(http.StatusInternalServerError, map[string]string{
					"mobile": "登陆失败",
				})
			}
			return
		}
	} else {
		//查需要调用底层才能校验密码。去python那边定义接口
		//	之所以先查用户名，是因为更通用。其实两次查询在性能上，是不如一次认证的。
		zap.S().Info(rsp.Password)
		passRsp, passErr := global.UserSrvClient.CheckPassword(context.Background(), &proto.PasswordCheck{
			Password:  passwordLoginForm.Password, //请求中的密码
			Encrypted: rsp.Password,               //询到用户后，其实已经可以用rsp.Password了等于加密后的密码
		})
		if passErr != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{
				"password": "登录失败",
			})
		} else {
			if passRsp.Success {
				j := middlewares.NewJWT()
				claims := models.CustomClaims{
					ID:          uint(rsp.Id),
					NickName:    rsp.NickName,
					AuthorityId: uint(rsp.Role),
					StandardClaims: jwt.StandardClaims{
						NotBefore: time.Now().Unix(),
						ExpiresAt: time.Now().Unix() + 60*60*24*30,
						Issuer:    "alice",
					},
				}
				token, err := j.CreateToken(claims)
				if err != nil {
					c.JSON(http.StatusBadRequest, map[string]string{
						"msg": "登陆失败",
					})
					return
				}

				c.JSON(http.StatusOK, gin.H{
					"id":         rsp.Id,
					"nick_name":  rsp.NickName,
					"token":      token,
					"expired_at": (time.Now().Unix() + 60*60*24*30) * 1000,
				})

				//c.JSON(http.StatusOK, map[string]string{
				//	"msg": "登陆成功",
				//})
			} else {
				c.JSON(http.StatusOK, map[string]string{
					"msg": "登录失败",
				})

			}
		}
	}
}

func Register(c *gin.Context) {
	//表单验证，验证码验证，新建用户，返回结果（登录的token）
	registerForm :=forms.RegisterForm{}
	if err:=c.ShouldBind(&registerForm);err!=nil{
		HandlerValidatorError(c,err)
		return
	}
	//验证码
	rdb:=redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",global.ServerConfig.RedisInfo.Host,global.ServerConfig.RedisInfo.Port),
	})
	value,err:=rdb.Get(registerForm.Mobile).Result()
	if err ==redis.Nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"msg":"验证码错误",
		})
		return
	}else {
		if value!=registerForm.Code{
			c.JSON(http.StatusBadRequest,gin.H{
				"msg":"验证码错误",
			})
			return
		}
	}
	//新建用户
	//userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", global.ServerConfig.UserSrvConfig.Host,
	//	global.ServerConfig.UserSrvConfig.Port), grpc.WithInsecure())
	//if err != nil {
	//	zap.S().Errorw("[GetUserList]连接 【用户服务】失败", "msg", err.Error())
	//}
	//userSrvClient := proto.NewUserClient(userConn)
	//以上从global获取

	user,err:=global.UserSrvClient.CreateUser(context.Background(),&proto.CreateUserInfo{
		NickName: registerForm.Mobile,
		Password: registerForm.Password,
		Mobile: registerForm.Mobile,
	})
	if err!=nil{
		zap.S().Errorf("[Register]查询 【新建用户失败】：%s",err.Error())
		HandlerGrpcErrorToHttp(err,c)
		return
	}

	j:=middlewares.NewJWT()
	claims := models.CustomClaims{
		ID: uint(user.Id),
		NickName: user.NickName,
		AuthorityId: uint(user.Role),
		StandardClaims:jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix()+60*60*24*30,
			Issuer: "alice",
		},
	}

	token,err:=j.CreateToken(claims)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"msg":"生成token失败",
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"id":user.Id,
		"nick_name":user.NickName,
		"token":token,
		"expired_at":(time.Now().Unix()+60*60*24*30)*1000,
	})
}
