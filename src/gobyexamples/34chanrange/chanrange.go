package main

import "fmt"

/*
chan 的for range循环
关闭通道还可以从里面接收数据
 */

func main() {
	ch:=make(chan string,3)
	ch<- "hello"
	ch <- "lisi"
	close(ch)

	for v:=range ch{
		fmt.Println(v)
	}
}