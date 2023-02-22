package main

import "fmt"

/*
主协程和子协程通过管道进行传递数据
主协程没有要传递的，就关闭协程
子协程是个死循环，循环从jobs里面接收，
如果协程关闭，返回值more会是false .
当more为false时，通过done向主协程通知done

 */
func main() {
	jobs:=make(chan int,5)
	done:=make(chan bool)


	go func() {
		for  {
			j,more:=<-jobs
			if more{
				fmt.Println("receiving :",j)
			}else {
				fmt.Println("completed")
				done <- true
				return
			}
		}
	}()

	for j:=1;j<=3;j++{
		jobs <-j
		fmt.Println("sending job",j)
	}
	close(jobs)
	fmt.Println("sednding all")
	<-done
}
