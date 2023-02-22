package main

import (
	"fmt"
	"time"
)

/*
利用chan的阻塞， 等待协程完成
 */

func foo(done chan bool) {
	fmt.Println("working ")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true//用管道来通知已经完成
}

func main() {
	ch:=make(chan  bool,1)//开启一个管道
	foo(ch)

	<-ch//阻塞直到收到通知
}
