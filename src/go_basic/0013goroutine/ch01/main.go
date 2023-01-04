package main

import (
	"fmt"
	"time"
)

func p() {
	fmt.Println("alice goroutine")
}
func loopP() {
	for {
		fmt.Printf("goroutine 1 \n")
	}
}

func main() {
	//go p()
	//go loopP()

	//匿名函数写法
	//go func() {
	//	for  {
	//		fmt.Printf("goroutine 2 \n")
	//	}
	//}()
	//主死从随。主进程\主线程 运行完，协程可能不运行

	//开启多个协程n
	for i:=0;i<5;i++{
		go func(n int) {
			for  {
				fmt.Printf("goroutine %d \n",n)//协程怎么使用到主协程的变量i
				time.Sleep(time.Second)
			}
		}(i)
	}
	time.Sleep(time.Second * 2)
	fmt.Printf("alice major")
}
