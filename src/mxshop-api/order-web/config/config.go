package config

type GoodsSrvConfig struct {
	//用condul，不需要再定义host  port
	//Host string `mapstructure:"host"`
	//Port int    `mapstructure:"port"`
	Name string `mapstructure:"name"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key"`
}

type ServerConfig struct {
	Name           string         `mapstructure:"name"` //mapstruture缺少c
	Port           int            `mapstructure:"port"`
	Host           string         `mapstructure:"host"`
	Tags           []string       `mapstructure:"tags"`
	GoodsSrvConfig GoodsSrvConfig `mapstructure:"goods_srv" json:"goods_srv"`
	OrderSrvInfo   GoodsSrvConfig `mapstructure:"order_srv" json:"order_srv"`//复用上面的类型
	JWTInfo        JWTConfig      `mapstructure:"jwt" json:"jwt"`
	ConsulInfo     ConsulConfig   `mapstructure:"consul" json:"consul"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}


type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      uint64    `mapstructure:"port"`
	Namespace string `mapstructure:"namespace"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	DataId    string `mapstructure:"dataid"`
	Group     string `mapstructure:"group"`
}