package main

import (
	"fmt"
	"time"
)

var timeout bool =false


func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch1 <- 1
	ch2 <- 2

	select {
	case data := <-ch1:
		fmt.Println(data)
	case data := <-ch2:
		fmt.Println(data)
	}//循环公平进入

	//
	ch3 :=make(chan bool,1)
	go func() {
		time.Sleep(time.Second*2)
		ch3<- true
	}()
	fmt.Println(<-ch3)

	ch4:=make(chan bool,1)
	ch5:=make(chan bool,1)
	go func() {
		time.Sleep(time.Second*2)
		ch4<- true
	}()
	go func() {
		time.Sleep(time.Second*4)
		ch5<- true
	}()

	for i:=0;i<10;i++{
		select {
		case <-ch4:
			fmt.Println("ch4取到值了")
		case <-ch5:
			fmt.Println("ch5取到值了")
			//break
		default:
			fmt.Println("都未取到值")
			time.Sleep(time.Second)
		}

	}


}
