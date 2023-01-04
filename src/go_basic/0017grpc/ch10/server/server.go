package main

import (
	context "context"
	"fmt"
	"go_learning/src/go_basic/0017grpc/ch10/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net"
)

type Server struct {
}

func (s Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	md ,ok:=metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println(ok)
		fmt.Println("get mettadtat error")
	}
	for k,v :=range md{
		fmt.Println(k,v)
	}
	if nameSlice,ok := md["auth"];ok{
		fmt.Println(nameSlice)
		for i,v :=range nameSlice{
			fmt.Println(i,v)
		}
	}
	return &proto.HelloReply{Message: "hello "+request.Name},nil
	//panic("implement me")
}

func main() {
	g := grpc.NewServer()
	proto.RegisterGreeterServer(g,&Server{})

	lis,err :=net.Listen("tcp","0.0.0.0:8889")
	if err!=nil {
		panic("failed to listen ：" + err.Error())
	}
	err=g.Serve(lis)
	if err!=nil{
		panic("failed to start grpc : "+err.Error())
	}


}
