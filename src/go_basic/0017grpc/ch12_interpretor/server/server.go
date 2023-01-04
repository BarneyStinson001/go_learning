package main

import (
	context "context"
	"fmt"
	"go_learning/src/go_basic/0017grpc/ch10/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net"
	"time"
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
	interceptor :=  func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error){
		recv_time := time.Now()
		fmt.Println(recv_time)
		fmt.Println("收到请求")
		res,err:= handler(ctx,req)
		fmt.Println(time.Since(recv_time))
		return res,err
	}
	opt:=grpc.UnaryInterceptor(interceptor)//而这个的入参为UnaryServerInterceptor。
	g := grpc.NewServer(opt)//可选参数：ServerOption，可以用UnaryInterceptor生成（返回类型为ServerOption）
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
