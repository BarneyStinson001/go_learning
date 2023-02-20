package main

import "fmt"

func plus(a, b int) int {
	return a+b
}

func plusPlus(a,b,c int) int	  {
	return a+b+c
}

func main() {
	fmt.Println(plus(1,2))
	fmt.Println(plusPlus(1,2,3))

	//l:=[]int{1,2,3}
	//plusPlus()

}