package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"time"
)

func main() {
	c,err:=rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{"192.168.18.160:9876"}),
		consumer.WithGroupName("qwee"),//group的作用
		)
	if err!=nil{
		panic("启动消费者失败")
	}
	err = c.Subscribe("test",consumer.MessageSelector{}, func(ctx context.Context, ext ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for i :=range ext{
			fmt.Printf("获取到值%v \n",ext[i])
		}
		return consumer.ConsumeSuccess,nil
	})
	if err!=nil{
		fmt.Println("读取消息失败")
	}

	_ = c.Start()
	time.Sleep(time.Hour)
	_=c.Shutdown()
}