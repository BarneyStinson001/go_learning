package main

import (
	"context"
	"fmt"
	"go_learning/src/go_basic/0017grpc/ch10/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	conn,err := grpc.Dial("127.0.0.1:8889",grpc.WithInsecure())
	if err!=nil {
		panic(err)
	}
	defer conn.Close()


	c:=proto.NewGreeterClient(conn)
	//r,err:=c.SayHello(context.Background(),&proto.HelloRequest{Name: "alice"})

	//带上metadata，先实例化，再带上
	md :=metadata.Pairs("auth","abcnsdckjdkcld")
	ctx := metadata.NewOutgoingContext(context.Background(),md)
	r,err:=c.SayHello(ctx,&proto.HelloRequest{Name: "alice"})


	if err!=nil{
		panic(err)
	}
	fmt.Println(r.Message)
}