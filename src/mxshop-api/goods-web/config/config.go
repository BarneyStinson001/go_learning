package config

type GoodsSrvConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
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
	GoodsSrvConfig GoodsSrvConfig `mapstructure:"goods_srv"`
	JWTInfo        JWTConfig      `mapstructure:"jwt"`
	ConsulInfo     ConsulConfig   `mapstructure:"consul"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
