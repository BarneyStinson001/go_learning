package main

import (
	"fmt"
	"time"
)

func main() {
	t1:=time.NewTicker(time.Second)
	done:=make(chan bool)


	go func() {
		for {//无限循环的协程
			select {//多路复用
			case <-done://通过done这个chan来通知这个协程结束
				return
			case t:=<-t1.C://每隔t就会有数据发到chan,这里就能执行
				fmt.Println("Tick at",t)
			}
		}
	}()
	time.Sleep(time.Millisecond*3500)
	t1.Stop()
	done <- true
	fmt.Println("tikcer stopped")
}
