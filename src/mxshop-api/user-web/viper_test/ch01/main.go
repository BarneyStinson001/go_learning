package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Name string `mapstructure:"name"`//为什么不叫yaml,因为不光支持yaml，还有json等
	Port int `mapstructure:"port"`
}
func main()	  {
	v:=viper.New()
	v.SetConfigFile("user-web/viper_test/ch01/config-debug.yaml")
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

}