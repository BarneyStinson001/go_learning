package main

import (
	"fmt"
	"time"
)

/*
在未来某个时间点执行，或者是重复间隔执行
内置timer  和ticker


 */

func main() {
	//定义2s倒计时，自带管道
	t1:=time.NewTimer(2*time.Second)
	//管道阻塞，直到倒计时归0,其中的chan有数据才能发
	<-t1.C
	fmt.Println("t1 over ")

	//相比于直接使用sleep,timer优势在于可以取消
	t2:=time.NewTimer(2*time.Second)
	go func() {
		//这个协程在等待倒计时
		<-t2.C
		fmt.Println("t2 over")
	}()
	//但主线程调用了stop
	stop:=t2.Stop()
	if stop{
		fmt.Println("t2 stopped")
	}
	time.Sleep(time.Second*2)


}