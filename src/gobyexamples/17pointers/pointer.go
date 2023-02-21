package main

import "fmt"

func zeroVal(n int)  {
	n = 0
}

func zeroPtr(n *int)  {
	*n=0
}

func main() {
	i:=1

	fmt.Println("init: ",i)
	zeroVal(i)
	fmt.Println("zeroVal: ",i)
	zeroPtr(&i)
	fmt.Println("zeroPtr: ",i)

	fmt.Println("pointer: ",&i)


}

/*
go语言允许进行指针操作

zeroVal  入参不是指针变量，是值传递。函数内操作不影响传入的变量

zeroPtr  参数*int  int型指针。  *n  传入地址，对该该地址的对应的值进行修改，影响到传入变量的值

 */