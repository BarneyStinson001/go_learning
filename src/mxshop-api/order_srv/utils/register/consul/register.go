package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

type Registry struct {
	Host string
	Port int
}

type RegisterClient interface {

	Register(address string,port int,name string,tags []string,id string) error
	Deregister(serviceId string) error
}

func NewRegistryClient(host string,port int)RegisterClient  {
	return &Registry{
		Host: host,
		Port: port,
	}
}

func (r *Registry)Register(address string,port int,name string,tags []string,id string) error {
	cfg:=api.DefaultConfig()
	cfg.Address=fmt.Sprintf("%s:%d",r.Host,r.Port)

	client,err:=api.NewClient(cfg)
	if err!=nil{
		panic(err)
	}

	//grpc的检查对象
	check:=&api.AgentServiceCheck{
		GRPC: fmt.Sprintf("%s:%d",address,port),
		Timeout: "5s",
		Interval: "5s",
		DeregisterCriticalServiceAfter: "15s",
	}

	//生成注册对象
	registration:=new(api.AgentServiceRegistration)
	registration.Name=name
	registration.ID=id
	registration.Port=port
	registration.Address=address
	registration.Tags=tags
	registration.Check=check

	err = client.Agent().ServiceRegister(registration)

	if err!=nil{
		panic(err)
	}
	return nil
}

func (r *Registry)Deregister(serviceId string) error {
	cfg:=api.DefaultConfig()
	cfg.Address=fmt.Sprintf("%s:%d",r.Host,r.Port)

	client,err:=api.NewClient(cfg)
	if err!=nil{
		return err
	}
	err= client.Agent().ServiceDeregister(serviceId)
	return err
}
