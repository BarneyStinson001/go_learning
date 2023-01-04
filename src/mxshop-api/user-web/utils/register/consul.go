package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

//尽管只是做接入consul进行服务注册和发现，实现过程可适当考虑扩展性

//1、读取consul的配置  如何初始化？

//
type Registery struct {
	Host string
	Port int
}

type RegisteryClient interface {
	Register(address string, port int, name string, tags []string, id string) error
	Deregister(id string)error
}

func NewRegisteryClienr(host string,port int)RegisteryClient{
	return &Registery{host,port}
}


func (r *Registery)Register(address string, port int, name string, tags []string, id string) error {
	cfg := api.DefaultConfig()
	//cfg.Address = "192.168.18.160:8500"
    //cfg.Address = fmt.Sprintf("%s:%d",address,port)  不是web服务器ip port
	cfg.Address = fmt.Sprintf("%s:%d",r.Host,r.Port)  //是consul服务器ip port

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	//生成注册对相关
	registration := new(api.AgentServiceRegistration)
	registration.Name = name
	registration.ID = id
	registration.Port = port
	registration.Tags = tags
	registration.Address = address

	//生成check
	check := &api.AgentServiceCheck{
		HTTP:                           fmt.Sprintf("http://%s:%d/health", address, port),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "5s",
	}
	registration.Check = check

	err = client.Agent().ServiceRegister(registration) //需要一个AgentServiceRegistration对象
	if err != nil {
		panic(err)
	}
	return nil //正常返回为空

}

func (r *Registery)Deregister(id string)error  {
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d",r.Host,r.Port)

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	err = client.Agent().ServiceDeregister(id)
	if err!=nil{
		panic(err)
	}
	return nil
}
