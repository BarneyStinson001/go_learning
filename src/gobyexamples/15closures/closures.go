package main

import "fmt"

func NextInt() func() int  {
	i:=0
	return func() int {
		i++
		return i
	}
}

//返回的是个函数，函数是匿名定义的，变量i是内部定义的。
//新定义的函数，都有自己的i值
func main() {
	next:=NextInt()

	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())

	next2:=NextInt()
	fmt.Println(next2())
	fmt.Println(next2())

}
