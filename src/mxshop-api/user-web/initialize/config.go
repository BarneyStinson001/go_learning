package initialize

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"mxshop-api/user-web/global"
)

func GetEnvInfo(env string)bool{
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {
	debug:= GetEnvInfo("msshop_debug")
	configName:="user-web/config-pro.yaml"
	if debug{
		fmt.Println("读入debug配置文件")
		configName="user-web/config-debug.yaml"
	}
	v:=viper.New()
	v.SetConfigFile(configName)
	if err:= v.ReadInConfig();err!=nil{
		panic(err)
	}
	value:=v.Get("name")
	fmt.Println(value)

	//serverConfig:=config.ServerConfig{}//这个对象如何在其他文件中使用 ，需要是全局变量。在global那边初始化，这边涉略初始化
	if err:=v.Unmarshal(global.ServerConfig);err!= nil {
		panic(err)
	}
	//fmt.Println(global.ServerConfig)
	zap.S().Infof("配置信息：%v",global.ServerConfig)//注意占位符使用
	fmt.Println(GetEnvInfo("msshop_debug"))

	fmt.Println("配置监控")
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		//fmt.Println("config file changed: ",in.Name)
		zap.S().Infof("配置文件产生变化： %s",in.Name)
		_=v.ReadInConfig()
		_=v.Unmarshal(global.ServerConfig)
		fmt.Println(global.ServerConfig)
	}) 

}