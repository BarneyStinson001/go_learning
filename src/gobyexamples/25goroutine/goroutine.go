package main

import (
	"fmt"
	"time"
)

/*
轻量级线程，即协程
函数调用方法：
协程调用函数：
协程调用匿名函数：

为了避免主死随从  time.sleep
 */

func foo(from string)  {
	for i:=0;i<3;i++{
		time.Sleep(time.Second)
		fmt.Println(from,":",i)
	}
}

func main() {
	foo("direct")

	go foo("goroutine")
	go foo("goroutine2")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(5*time.Second)
	fmt.Println("done")
}
