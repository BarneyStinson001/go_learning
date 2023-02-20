package main

import "fmt"

func sum(nums ...int	)int  {
	fmt.Print(nums," ")
	t:=0
	for _,num :=range  nums{
		t+=num
	}

	fmt.Println(t)
	return t
}

func main()  {
	sum(1,2,3)
	sum(1)

	nums:=[]int{1,2,3,4,5,6}
	sum(nums...)
}
