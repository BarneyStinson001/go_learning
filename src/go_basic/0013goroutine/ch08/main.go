package main

import (
	"fmt"
	"sync"
	"time"
)

//发送方关闭
//消费者用for range读就没问题了
//
var wg sync.WaitGroup
//
//func consumer(queue chan int) {
//	defer wg.Done()
//	for {
//		data, ok := <-queue //从管道获取data,判断是否已经关闭
//		if !ok {
//			break
//		}
//		fmt.Println(data)
//		time.Sleep(time.Second)
//	}
//}
//func main() {
	//var msg chan int
	//msg = make(chan int)
	//wg.Add(1)
	//go consumer(msg)
	//msg <- 1
	//msg <- 2
	//msg <- 3
	//close(msg) //关闭的管道不会再发送数据
	//wg.Wait()
//}

//上面为无缓冲的chan
//13-08 chan类型


func consumer(queue chan int) {
	defer wg.Done()
	for {
		data, ok := <-queue //从管道获取data,判断是否已经关闭
		if !ok {
			break
		}
		fmt.Println(data)
		time.Sleep(time.Second)
		queue <- 2 //返回2
	}
}
func main() {
	var msg chan int
	msg = make(chan int)
	wg.Add(1)
	go consumer(msg)
	msg <- 1
	fmt.Println("waiting")
	fmt.Println(<- msg)//拿到返回的2
	close(msg)
	wg.Wait()


	//单向chan   var msg chan<- int   只能放进来
	//单向chan   var msg <-chan int   只能取出去
	//一般使用：func consumer(queue <-chan int) {
	//主的是chan双向，还是可以放，但consumer只能取
	//双向chan直接转换为单向管道

	//爬虫 主goroutine发现url，让子goroutine去拉，但不用产生
	//内存监控，子goroutine放数据，主取数据分析
}