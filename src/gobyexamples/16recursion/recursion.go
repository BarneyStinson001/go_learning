package main

import "fmt"

/*
经典递归：
斐波那契数列
阶乘
-----尽管不是一个很高效的算法
 */

func fact(n int) int  {
	if n==0{
		return 1
	}
	return n*fact(n-1)
}

func main() {
	fmt.Println(fact(5))


	var fib func(n int)int
	fib= func(n int) int {
		if n<2 {
			return n
		}
		return fib(n-1)+fib(n-2)
	}

	fmt.Println(fib(7))
}