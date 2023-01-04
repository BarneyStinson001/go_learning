package main

import (
	"context"
	"fmt"
	"go_learning/src/go_basic/0016grpc/ch08/proto"
	"google.golang.org/grpc"
	"sync"
	"time"
)

func main() {
	conn ,err := grpc.Dial("localhost:50052",grpc.WithInsecure())
	if err!=nil{
		panic(err)
	}

	defer conn.Close()

	c:=proto.NewGreeterClient(conn)

	//服务端响应流模式
	res,_:=c.GetStream(context.Background(),&proto.StreamReqData{Data: "Alice"})
	for{
		//一recv就打印
		a,err :=res.Recv()//socket编程
		if err != nil{
			fmt.Println(err)
			break
		}
		fmt.Println(a)
	}

	//客户端发送流模式，server有一直收到
	putS,_:=c.PostStream(context.Background())

	i:=0
	for {
		i++
		putS.Send(&proto.StreamReqData{
			Data: fmt.Sprintf("alice %d",i),
		})
		time.Sleep(time.Second)
		if i >2{
			break
		}
	}

	//双向流模式
	allStream,_:=c.AllStream(context.Background())
	wg:=sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for{
		msg,_:=allStream.Recv()
		fmt.Println("收到服务端信息："+msg.Data)}
	}()

	go func() {
		defer wg.Done()
		for{
			_=allStream.Send(&proto.StreamReqData{Data:"client req"})
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()

}
