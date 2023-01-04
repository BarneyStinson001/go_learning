package initialize

import (
	"encoding/json"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"mxshop-api/order-web/global"
)

func GetEnvInfo(env string)bool{
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {
	debug:= GetEnvInfo("msshop_debug")
	zap.S().Info(GetEnvInfo("msshop_debug"))

	configName:="order-web/config-pro.yaml"
	if debug{
		zap.S().Info("使用debug配置文件，非production环境")
		configName="order-web/config-debug.yaml"
	}
	v:=viper.New()
	v.SetConfigFile(configName)
	if err:= v.ReadInConfig();err!=nil{
		panic(err)
	}
	value:=v.Get("name")
	zap.S().Info(value)

	//serverConfig:=config.ServerConfig{}//这个对象如何在其他文件中使用 ，需要是全局变量。在global那边初始化，这边涉略初始化
	if err:=v.Unmarshal(&global.NacosConfig);err!= nil {
		panic(err)
	}
	//fmt.Println(global.ServerConfig)
	zap.S().Infof("配置信息：%v",global.NacosConfig)//注意占位符使用

	//fmt.Println("配置监控")
	//v.WatchConfig()
	//v.OnConfigChange(func(in fsnotify.Event) {
	//	//fmt.Println("config file changed: ",in.Name)
	//	zap.S().Infof("配置文件产生变化： %s",in.Name)
	//	_=v.ReadInConfig()
	//	_=v.Unmarshal(global.ServerConfig)
	//	fmt.Println(global.ServerConfig)
	//})

	//从nacos中读取配置信息
	sc := []constant.ServerConfig{
		{
			IpAddr: global.NacosConfig.Host,
			Port: global.NacosConfig.Port,
		},
	}

	cc := constant.ClientConfig {
		NamespaceId:         global.NacosConfig.Namespace, // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		//RotateTime:          "1h",
		//MaxAge:              3,
		LogLevel:            "debug",
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		panic(err)
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: global.NacosConfig.DataId,
		Group:  global.NacosConfig.Group})

	if err != nil {
		panic(err)
	}
	zap.S().Info(content)
	//fmt.Println(content) //字符串 - yaml
	//想要将一个json字符串转换成struct，需要去设置这个struct的tag
	err = json.Unmarshal([]byte(content), &global.ServerConfig)
	if err != nil{
		zap.S().Fatalf("读取nacos配置失败： %s", err.Error())
	}
	zap.S().Info(global.ServerConfig)
	//&{order-web 8022 192.168.18.179 [mxshop order web] {} {} {} { 0}}
	//{orders-srv 192.168.18.179 [gorm srv order] {192.168.18.160 3306 mxshop_order_srv root root} {192.168.18.160 8500} { 0} {goods_srv} {invenotry-srv}}


}