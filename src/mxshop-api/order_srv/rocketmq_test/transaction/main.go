package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"time"
)

type OrderListener struct {
	////  When send transactional prepare(half) message succeed, this method will be invoked to execute local transaction.
	//ExecuteLocalTransaction(*Message) LocalTransactionState
	//
	//// When no response to prepare(half) message. broker will send check message to check the transaction status, and this
	//// method will be invoked to get local transaction status.
	//CheckLocalTransaction(*MessageExt) LocalTransactionState
}

func (o *OrderListener)ExecuteLocalTransaction(msg *primitive.Message) (primitive.LocalTransactionState) {
	return primitive.CommitMessageState
}

func (o *OrderListener)CheckLocalTransaction(ext *primitive.MessageExt)(primitive.LocalTransactionState)  {
	return primitive.RollbackMessageState
}


func main()	  {
	//trtansactionlistener 事务监听，回查事务
	p,err:=rocketmq.NewTransactionProducer(
		&OrderListener{},
		producer.WithNameServer([]string{"192.168.18.160:9876"}))
	if err!=nil{
		panic("生成producer失败")
	}

	if err=p.Start();err!= nil {
		panic("启动producer失败")
	}

	res,err:=p.SendMessageInTransaction(context.Background(),primitive.NewMessage("trans",[]byte("transaction msg")))
	if err!=nil{
		fmt.Printf("发送失败%s\n",err)
	}else{
		fmt.Printf("发送成功%是\n",res.String())
	}


	time.Sleep(time.Hour)
	if err= p.Shutdown();err!=nil{
		panic("关闭producer失败")
	}


}
