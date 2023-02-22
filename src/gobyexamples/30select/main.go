package main

import (
	"fmt"
	"time"
)

/*
多个协程并发执行，同时监听，谁先结束，就进行响应处理
 */

func main() {
	c1:= make(chan string)
	c2:=make(chan string)

	go func() {
		time.Sleep(time.Second)
		c2<-"two"
	}()

	go func() {
		time.Sleep(time.Second*2)
		c1<-"one"
	}()

	for i:=0;i<2;i++{
		select {
		case msg := <-c1:
			fmt.Println("received from c1",msg)
		case msg2 := <-c2:
			fmt.Println("received from c2",msg2)
		}

	}
}