package main

import "fmt"

func main() {
	nums:=[]int{1,2,3,4,5}
	sum:=0
	for _,val:=range nums{
		sum+=val
	}
	fmt.Println("sum: ",sum)

	for i,num:=range nums{
		fmt.Println(i,num)
	}

	m:=map[string]int{"lisi":90,"zhangsan":80}
	for k,v:=range m{
		fmt.Printf("%s  -> %d",k,v)
	}
	fmt.Println()

	for k:=range m{
		fmt.Println("key: ",k)
	}

	for i,c:=range "abnsdfdf"{
		fmt.Println(i,c)
	}

}


/*
迭代：切片、map、字符串

 */