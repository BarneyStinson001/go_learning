package main

import (
	"context"
	"fmt"
	"go_learning/src/go_basic/0017grpc/ch10/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"
)

func main() {
	interceptor :=func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error{
		start := time.Now()
		err :=invoker(ctx,method,req,reply,cc,opts...)
		fmt.Println("耗时： %s",time.Since(start))
		return err
	}
	diaoption := grpc.WithUnaryInterceptor(interceptor)
	//conn,err := grpc.Dial("127.0.0.1:8889",grpc.WithInsecure(),diaoption)
	//第二种：
	var diaoptions []grpc.DialOption
	diaoptions =append(diaoptions,grpc.WithInsecure(),)
	diaoptions =append(diaoptions,diaoption)
	conn,err := grpc.Dial("127.0.0.1:8889",diaoptions...)
	if err!=nil {
		panic(err)
	}
	defer conn.Close()


	c:=proto.NewGreeterClient(conn)
	//r,err:=c.SayHello(context.Background(),&proto.HelloRequest{Name: "alice"})

	//带上metadata，先实例化，再带上
	md :=metadata.Pairs("auth","wodemima")
	ctx := metadata.NewOutgoingContext(context.Background(),md)
	r,err:=c.SayHello(ctx,&proto.HelloRequest{Name: "linda"})


	if err!=nil{
		panic(err)
	}
	fmt.Println(r.Message)
}