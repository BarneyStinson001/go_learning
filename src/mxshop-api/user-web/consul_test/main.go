package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

func Register(address string, port int, name string, tags []string, id string) error {
	cfg := api.DefaultConfig()
	cfg.Address = "192.168.18.160:8500"

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
func Deregister(id string)error  {
	cfg := api.DefaultConfig()
	cfg.Address = "192.168.18.160:8500"

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

func AllService() {
	cfg := api.DefaultConfig()
	cfg.Address = "192.168.18.160:8500"

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	data, err := client.Agent().Services()
	if err != nil {
		panic(err)
	}

	for key, value := range data {
		fmt.Printf("key: %s\n", key)
		fmt.Printf("value: %T\n", value)

	}
}

func FilterServive(name string) {
	cfg := api.DefaultConfig()
	cfg.Address = "192.168.18.160:8500"

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	data, err := client.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%s"`, name))
	if err != nil {
		panic(err)
	}

	for key, value := range data {
		fmt.Printf("key: %s, value: %v\n", key, value)
	}
}

func RegisterGRPCService(address string, port int, name string, tags []string, id string) error {
	cfg := api.DefaultConfig()
	cfg.Address = "192.168.18.160:8500"

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

	//生成checkf
	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d", address, port),
		GRPCUseTLS:                     false,
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

func main() {
	//_ = Register("192.168.18.179", 8021, "mxshop-web", []string{"web", "mxshop-web"}, "mxshop-web")
	//if err!=nil{
	//	panic(err)
	//}
	//_ = RegisterGRPCService("192.168.18.179", 50021, "mxshop-srv", []string{"service", "mxshop-srv"}, "mxshop-srv")
	fmt.Printf("=========list\n")
	AllService()
	fmt.Printf("=========Filter\n")
	FilterServive("user-srv")

	//Deregister("mxshop-web")

}
