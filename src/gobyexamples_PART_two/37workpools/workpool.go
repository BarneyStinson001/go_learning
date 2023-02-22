package main

import (
	"fmt"
	"time"
)

/*
工作池

 */

func worker(id int,jobs <-chan int,results chan<- int)  {
	for j:=range  jobs{
		fmt.Println("worker ",id , "start job",j)
		time.Sleep(time.Second)
		fmt.Println("worker ",id , "finished job",j)
		results <-j *2
	}
}

func main() {
	const num = 5

	jobs:=make(chan  int,num)
	res:=make(chan int,num)

	for i:=1;i<=3;i++{
		go worker(i,jobs,res)
	}

	for j:=1;j<=num;j++{
		jobs <- j
	}
	close(jobs)

	for a:=0;a<num;a++{
		fmt.Println(<- res)
	}
}
