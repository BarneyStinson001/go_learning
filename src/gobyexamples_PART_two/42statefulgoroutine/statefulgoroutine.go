package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

/*
除了用互斥锁，显式的锁住数据，来同步访问数据，
管道也可以实现，和go的设计（通过通信来共享内存）是一致的，
数据只由一个协程拥有


读写状态由单独的协程所有，
保证在并发访问过程中不会出现数据竞争。
为了进行读或写操作，其他协程会发送信息给owner.以及收到响应。
结构体封装了响应体。


*/

type readOp struct {
	key int
	rsp chan int
}

type writeOP struct {
	key int
	val int
	rsp chan bool
}

func main() {
	var readOps uint64
	var writeOps uint64

	reads := make(chan readOp) //无缓冲队列，可能发生阻塞
	writes := make(chan writeOP)

	//开一个协程，
	go func() {
		//是state的 owner
		var state = make(map[int]int)
		//死循环
		for {
			//当读和写管道有数据过来后，做出响应。
			select {
			case read := <-reads:
				read.rsp <- state[read.key]//传入read.key，得到对应的state,放到rsp
			case write := <-writes:
				state[write.key] = write.val//把k,v设置到state  ,并把rsp记为成功
				write.rsp <- true
			}
		}
	}()

	for r:=0;r<100;r++{
		//开启100个协程
		go func() {
			//死循环
			for  {
				//构造readOp结构体
				read:=readOp{
					key: rand.Intn(5),
					rsp: make(chan int),
				}
				//通过reads进行读操作
				reads<-read
				//接收响应结果
				<-read.rsp
				atomic.AddUint64(&readOps,1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w:=0;w<10;w++{
		go func() {
			for true {
				write:=writeOP{
					key: rand.Intn(5),
					val: rand.Intn(100),
					rsp: make(chan bool),
				}
				writes<-write
				<- write.rsp
				atomic.AddUint64(&writeOps,1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	//没有定义wg,协程里是死循环。这里就只跑一分钟
	time.Sleep(time.Second)
	//最终的操作数可能存在变化，用过atomic.LoadUint64 获得。原子操作
	readOpsFinal:=atomic.LoadUint64(&readOps)
	fmt.Println("readOps",readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps",writeOpsFinal)




}
