package main

import "fmt"

/*
channal 是并发协程之间的管道，
可以从一个协程把数据发到管道里
也可以在协程里从管道里接收数据

管道分类型 int string
不定义容量，是没有缓冲的。

<-chan 如果没有chan<-  会一直阻塞

 */

func main() {
	mq:=make(chan string)

	go func() {mq <- "sending"}()

	msg:= <-mq
	fmt.Println(msg)
}