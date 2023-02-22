package main

import (
	"fmt"
	"time"
)

func main() {
	c1:=make(chan string,1)

	go func() {
		time.Sleep(time.Second*2)
		c1<- "result1"
	}()

	select {
	case res:=<-c1:
		fmt.Println("got result: ",res)
	case <-time.After(1*time.Second):
		fmt.Println("timeout 1")
	}
	//select同时监听，并接收第一个收到的。当res需要大于1秒时，就走了第二个分支。
	//下面的例子，仅修改了等待时间，所以在第二个分支进去超前，也就是超时前，能从管道接收
	c2:=make(chan string,1)

	go func() {
		time.Sleep(time.Second*2)
		c2<- "result2"
	}()

	select {
	case res:=<-c2:
		fmt.Println("got result: ",res)
	case <-time.After(3*time.Second):
		fmt.Println("timeout 2")
	}
}
