package main

import (
	"context"
	"fmt"
	"go_learning/src/go_basic/0016grpc/ch05/grpc_test/proto"
	"google.golang.org/grpc"
)

func main() {
	conn,err :=grpc.Dial("127.0.0.1:8090",grpc.WithInsecure())
	if err !=nil{
		panic(err)
	}
	defer conn.Close()//关闭


	//实例化
	c:=proto.NewGreeterClient(conn)
	r,err :=c.SayHello(context.Background(),&proto.HelloRequest{Name: "Alice"})
	if err!=nil{
		panic(err)
	}
	fmt.Println(r.Message)


}
