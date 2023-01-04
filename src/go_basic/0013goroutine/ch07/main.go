package main

import (
	"fmt"
	"sync"
	"time"
)

//发送方关闭
//消费者用for range读就没问题了 

var wg sync.WaitGroup

func consumer(queue chan int) {
	defer wg.Done()
	for {
		data, ok := <-queue//从管道获取data,判断是否已经关闭
		if !ok{
			break
		}
		fmt.Println(data)
		time.Sleep(time.Second)
	}
}

func main() {
	var msg chan int
	msg = make(chan int,1)
	msg <- 1

	wg.Add(1)
	go consumer(msg)
	msg <- 2
	msg <- 3

	close(msg)//关闭的管道不会再发送数据
	wg.Wait()
}
