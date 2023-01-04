package main

import (
	"context"
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"google.golang.org/grpc"
	"log"
	"mxshop-api/user-web/proto"
)

func main() {
	conn, err := grpc.Dial(
		"consul://192.168.18.160:8500/user-srv?wait=14s&python=srv",
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`), //关于配置看grpc官方文档,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	//循环请求测试
	for i:=0;i<10;i++{
	userSrvClient := proto.NewUserClient(conn)
	rsp, err := userSrvClient.GetUserList(context.Background(), &proto.PageInfo{
		PageNo:   1,
		PageSize: 2,
	})
	if err != nil {
		panic(err)
	}

	for index, value := range rsp.Data {
		fmt.Println(index, value)
	}
	}

}
