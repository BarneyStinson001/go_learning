package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"time"
)



//struct嵌套
type MysqlConfig struct {
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
}
type ServerConfig struct {
	Name string `mapstructure:"name"`//为什么不叫yaml,因为不光支持yaml，还有json等
	Port int `mapstructure:"port"`
	MysqlInfo MysqlConfig `mapstructure:"mysql"`
}

func GetEnvInfo(env string)bool{
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func main()	  {
	debug:= GetEnvInfo("msshop_debug")
	configName:="user-web/viper_test/ch02/config-produ.yaml"
	if debug{
		configName="user-web/viper_test/ch02/config-debug.yaml"
	}
	v:=viper.New()
	v.SetConfigFile(configName)
	if err:= v.ReadInConfig();err!=nil{
		panic(err)
	}
	namevalue:=v.Get("name")
	fmt.Println(namevalue)

	serverConfig:=ServerConfig{}
	if err:=v.Unmarshal(&serverConfig);err!= nil {
		panic(err)
	}
	fmt.Println(serverConfig)

	fmt.Println(GetEnvInfo("msshop_debug"))

	//配置监控功能
	fmt.Println("配置监控")
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed: ",in.Name)
		_=v.ReadInConfig()
		_=v.Unmarshal(&serverConfig)
		fmt.Println(serverConfig)
	})
	time.Sleep(time.Second*300)
}