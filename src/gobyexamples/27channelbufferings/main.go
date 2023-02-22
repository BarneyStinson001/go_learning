package main

import "fmt"

/*
有缓冲chan
可以接收多个值，不大于容量
然后陆续取出
 */

func main	()  {
	ch:=make(chan string,2)
	ch <- "hello"
	ch <- "hi"

	fmt.Println(<-ch)
	fmt.Println(<-ch)

}
