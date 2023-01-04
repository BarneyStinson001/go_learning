package config

type UserSrvConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	Name string `mapstructure:"name"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key"`
}

type SMSConfig struct {
	Width int `mapstructure:"width"`
}

type RedisConfig struct {
	Host   string `mapstructure:"host"`
	Port   int    `mapstructure:"port"`
	Expire int    `mapstructure:"expire"`
}

type ServerConfig struct {
	Name          string        `mapstructure:"name"` //mapstruture缺少c
	Port          int           `mapstructure:"port"`
	UserSrvConfig UserSrvConfig `mapstructure:"user_srv"`
	JWTInfo       JWTConfig     `mapstructure:"jwt"`
	SMSInfo       SMSConfig     `mapstructure:"sms"`
	RedisInfo     RedisConfig   `mapstructure:"redis"`
	ConsulInfo    ConsulConfig  `mapstructure:"consul"`
	Host          string        `mapstructure:"host"`
	Tags          []string		`mapstructure:"tags"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
