package main

import "fmt"

/*

 */

//ping只能往pings管道里发msg
func ping(pings chan<- string,msg string)  {
	pings <- msg
}
//pongs只能从pings管道读,发给pongs管道
func pong(pings <-chan  string,pongs chan<- string)  {
	msg:=<-pings
	pongs <- msg
}

func main() {
	pings:=make(chan string,1)
	pongs:=make(chan string,1)
	ping(pings,"sending msg")
	pong(pings,pongs)
	fmt.Println(<-pongs)
}