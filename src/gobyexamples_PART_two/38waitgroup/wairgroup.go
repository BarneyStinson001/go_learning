package main

import (
	"fmt"
	"sync"
	"time"
)

/*
等待一组goroutine结束，可以用waitgroup

 */

func wokrer(id int) {
	fmt.Println("worker start:...",id)
	time.Sleep(time.Second)
	fmt.Println("worker end. ",id)
}

func main() {
	var wg sync.WaitGroup

	for i:=0;i<5;i++{
		wg.Add(1)
		i:=i//如果不重新定义i,把外面的i传入，会导致前面的work取到后面的id
		go func() {//匿名函数，用了闭包思想，真正的工作函数不需要知道wg
			defer wg.Done()
			wokrer(i)
		}()
	}

	wg.Wait()//直到count到0

}
