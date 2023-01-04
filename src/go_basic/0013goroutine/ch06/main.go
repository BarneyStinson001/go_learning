package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func consumer(queue chan int) {
	defer wg.Done()
	data := <- queue
	fmt.Println(data)
}

func main() {
	//定义channal
	var msg chan int
	//初始化
	//msg = make(chan int)//无缓冲
	msg = make(chan int,1)//带缓冲空间，小心缓冲空间满，满了会deadlock死锁
	msg <- 1
	wg.Add(1)
	go consumer(msg)
	wg.Wait()
}