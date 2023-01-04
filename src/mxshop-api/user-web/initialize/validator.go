package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"mxshop-api/user-web/global"
	"reflect"
	"strings"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/go-playground/validator/v10"
)



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
		global.Trans, ok = uni.GetTranslator(locale)//设置为全局变量 不用：
		if !ok{
			return fmt.Errorf("uni.GetTranslator(%s)",locale)
		}
		//开始注册
		switch locale{
		case "en":en_translations.RegisterDefaultTranslations(v,global.Trans)
		case "zh":zh_translations.RegisterDefaultTranslations(v,global.Trans)
		default:
			en_translations.RegisterDefaultTranslations(v,global.Trans)
		}
		return
	}
	return
}