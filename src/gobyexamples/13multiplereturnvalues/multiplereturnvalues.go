package main

import (
	"fmt"
	"math/rand"
)

func random()  (int,int){
	return rand.Int(),rand.Int()
}

func main() {
	fmt.Println(random())

	a,b:=random()
	fmt.Println(a)
	fmt.Println(b)

	_,c:=random()
	fmt.Println(c)
}



/*
获取返回值和错误信息
 */