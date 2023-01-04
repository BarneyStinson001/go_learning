package main

import (
	"fmt"
	"go_learning/src/go_basic/0016grpc/ch08/proto"
	"google.golang.org/grpc"
	"net"
	"sync"
	"time"
)

const PORT = ":50052"

type server struct {

}
//服务端源源不断地返回时间戳，最后EOF
func (s *server)GetStream(data *proto.StreamReqData,res proto.Greeter_GetStreamServer)error  {
	i :=0
	for{
		i++
		_=res.Send(&proto.StreamResData{
			Data: fmt.Sprintf("%v",time.Now().Unix()),
		})
		time.Sleep( time.Second)
		if i>2 {
			break
		}
	}
	return nil//不能使用return
}
func (s *server)PostStream(data proto.Greeter_PostStreamServer)error {
	for{
		if a,err := data.Recv();err != nil  {
			fmt.Println(err)
			break
		}else{//如果收到EOF，就代表结束了，该响应了。
			fmt.Println(a.Data)
		}
	}
	return nil
}
func (s *server)AllStream(data proto.Greeter_AllStreamServer)error  {
	//不能allstr.Recv   allstr.Send需要并行，不能有前后关系。需要用协程
	wg:=sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for{
		msg,_:=data.Recv()
		fmt.Println("收到客户端信息："+msg.Data)}

	}()

	go func() {
		defer wg.Done()
		for{
			_=data.Send(&proto.StreamResData{Data:"server resp"})
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
	return nil
}

func main() {
	lis,err := net.Listen("tcp",PORT)
	if err!=nil {
		panic(err)
	}

	s:=grpc.NewServer()
	proto.RegisterGreeterServer(s,&server{})
	err = s.Serve(lis)
	if err!=nil {
		panic("failed to start grpc  :" +err.Error())
	}
}
