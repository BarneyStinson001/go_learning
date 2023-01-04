package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

func main() {
	p,err:=rocketmq.NewProducer(producer.WithNameServer([]string{"192.168.18.160:9876"}))
	if err!=nil{
		panic("生成producer失败")
	}
	
	if err=p.Start();err!= nil {
		panic("启动producer失败")
	}
	
	rsp,err:=p.SendSync(context.Background(),&primitive.Message{
		Topic:          "test",
		Body:           []byte("this is a test 2"),
		CompressedBody: nil,
		Flag:           0,
		TransactionId:  "",
		Batch:          false,
		Compress:       false,
		Queue:          nil,
	})

	if err!=nil{
		fmt.Printf("发送失败%s\n",err)
	}else{
		fmt.Printf("发送成功%是\n",rsp.String())
	}

	 if err= p.Shutdown();err!=nil{
	 	panic("关闭producer失败")
	 }
}