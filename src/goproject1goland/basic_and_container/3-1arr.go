package main

import "fmt"

func printArr(arr [5]int)  {//[5]int和【10】int不一样
	for i,v :=range arr{
		fmt.Println(i,v)
	}
	arr[0]=100//值传递，不修改到原数组
}
func printArr2(arr *[5]int)  {//传递指针
	for i,v :=range arr{
		fmt.Println(i,v)
	}
}



func main() {
	var arr [5]int
	arr2 := [3]int{1,3,5}
	arr3 := [...]int{2,4,6,8,10}
	var grid [4][5]int
	fmt.Println(arr,arr2,arr3,grid)

	for i:=0;i<len(arr3);i++{
		fmt.Println(arr3[i])
	}
	for i:=range arr3{
		fmt.Println(arr3[i])
	}
	//下标和值
	for i,v :=range arr3{
		fmt.Println(i,arr3[i],v)
	}
	//不用下标，就写下划线 
	for _,v :=range arr3{
		fmt.Println(v)
	}
	printArr(arr3)
	fmt.Println(arr3[0])

	printArr2(&arr3)
}

