package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"net/http"
	"reflect"
	"strings"
)

type LoginForm struct {
	User     string `json:"user" binding:"required,min=3,max=10"`
	Password string `json:"password" binding:"required"`
}

//注册表单号的验证
type SignUpForm struct {
	Age uint8 `json:"age" binding:"gte=1,lte=130"`
	Name string `json:"name" binding:"required,min=3"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	RePassword string `json:"repassword" binding:"required,eqfield=Password"`//注意格式正确，eqfield跨字段检查，

}
//全局变量
var trans ut.Translator


//22-10  {"error":{"LoginForm.Password":"Password为必填字段"}}
//---》"password":"password为必填字段"
//这里转化LoginForm.password 为 password
func removeTopStruct(fields map[string]string) map[string]string{
	rsp :=map[string]string{}
	for field,err:=range fields	{
		idx:=strings.Index(field,".")
		rsp[field[idx+1:]]=err
	}
	return rsp
}


func InitTrans(locale string)(err error)  {
	//修改gin框架中的validator属性，实现定制
	//
	if v,ok:=binding.Validator.Engine().(*validator.Validate);ok{
		//22-10把struct里面的key用json的tag返回回去
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"),",",2)[0]
			if name == "-"{
				return ""
			}
			return name
		})

		zhT :=zh.New()//中文翻译器
		enT :=en.New()
		uni :=ut.New(enT,zhT,enT)
		//trans, ok := uni.GetTranslator(locale)//用：会变成局部变量
		trans, ok = uni.GetTranslator(locale)//设置为全局变量 不用：
		if !ok{
			return fmt.Errorf("uni.GetTranslator(%s)",locale)
		}
	//开始注册
		switch locale{
		case "en":en_translations.RegisterDefaultTranslations(v,trans)
		case "zh":zh_translations.RegisterDefaultTranslations(v,trans)
		default:
			en_translations.RegisterDefaultTranslations(v,trans)
		}
		return
	}
	return
}


func main() {
	if err :=InitTrans("zh");err!=nil{
		fmt.Println("初始化出错")
		return
	}

	router := gin.Default()

	router.POST("/loginJSON", func(context *gin.Context) {
			var loginForm LoginForm
			if err:=context.ShouldBind(&loginForm);	err!=nil{
				errs,ok:=err.(validator.ValidationErrors)
				if !ok {
					context.JSON(http.StatusOK,gin.H{
						"msg":err.Error(),
					})
				}
 				fmt.Println(errs.Translate(trans))
				context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()},)
				context.JSON(http.StatusBadRequest,gin.H{
					"error":errs.Translate(trans),
				})
				//22-10 返回的值用removeTopStruct修饰下
				//context.JSON(http.StatusBadRequest,gin.H{
				//	"error":removeTopStruct(errs.Translate(trans)),
				//})
				return
			}
			context.JSON(http.StatusOK,gin.H{"msg":"登陆成功",})
	})
	router.POST("/signup", func(context *gin.Context) {
		var signUpForm SignUpForm
		if err:=context.ShouldBind(&signUpForm);err!=nil {
			fmt.Println(err.Error())
			context.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}
		context.JSON(http.StatusOK,gin.H{
			"msg":"register ok 注册成功",
		})
	})

	router.Run(":8083")
}
