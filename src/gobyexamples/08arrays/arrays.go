package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("empty: ",a)

	a[4]=100
	fmt.Println("set :",a)
	fmt.Println("get : ",a[4])

	fmt.Println(len(a))

	b:=[5]int{1,2,3,4,5}
	fmt.Println(b)

	var twoD [2][3]int
	for i:=0;i<2;i++{
		for j:=0;j<3;j++{
			twoD[i][j]=i+j
		}
	}
	fmt.Println("2d:",twoD)
}


/*
[5]int  [4]int  不是相同类型
go里面用的更多地是切片
零值可用，各类型的零值

设置值 a[i]=xx
随机访问 print(a[i])

len

命名并初始化

多维数组


 */