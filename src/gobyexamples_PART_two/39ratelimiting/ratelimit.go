package main

import (
	"fmt"
	"time"
)

/*
通过协程、管道、计时器，来控制速率

*/

func main() {
	reqs := make(chan int, 5)
	for i := 0; i < 5; i++ {
		reqs <- i
	}
	close(reqs) //传进来五个任务

	limiter := time.Tick(200 * time.Millisecond)//限制器：200ms

	for req := range reqs {
		<-limiter //每200ms通行一次
		fmt.Println("req ", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t//缓冲限制器通道：缓冲了3个
		}
	}()
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter //迅速有三个过来
		fmt.Println("request", req, time.Now())
	}

}
