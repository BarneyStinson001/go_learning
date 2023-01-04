package main

import (
	"context"
	"go_learning/src/go_basic/0016grpc/ch05/grpc_test/proto"
	"google.golang.org/grpc"
	"net"
)

type Server struct {

}
//	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
func (s *Server)SayHello(ctx context.Context,request *proto.HelloRequest)(*proto.HelloReply,error)  {
	return &proto.HelloReply{Message: "hello "+request.Name},nil
}

func main() {
	//实例化
	g := grpc.NewServer()
	//注册
	proto.RegisterGreeterServer(g,&Server{})//鸭子类型。用接口。不要绑定特定的
	//绑定和启动
	lis,err := net.Listen("tcp","0.0.0.0:8090")
	if err!=nil{
		panic("failed to listen :" +err.Error())
	}
	err = g.Serve(lis)
	if err!=nil{
		panic("failed to start grpc  :" +err.Error())
	}

}