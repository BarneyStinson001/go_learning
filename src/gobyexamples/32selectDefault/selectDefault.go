package main

import "fmt"

func main() {
	send:=make(chan<- string)
	recv:=make(<-chan string)
	recv2:=make(<-chan string)
	select {
	case msg := <-recv:
		fmt.Println("receiving:",msg)	
	default:
		fmt.Println("no receiving")
	}
	
	msg:="hello"
	select {
	case send <- msg:
		fmt.Println("sending :",msg)
	default:
		fmt.Println("no sending")
	}
	select {
	case rec:=<-recv:
		fmt.Println("receiving:",rec)
	case rec2:=<-recv2:
		fmt.Println("receiving:",rec2)
	default:
		fmt.Println("no activity")
	}

}
