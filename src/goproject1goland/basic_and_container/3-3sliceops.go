package main

import (
	"fmt"
)

func printArr(s []int)  {
	fmt.Printf("%v len id %d,cap is %d\n",s,len(s),cap(s))
}

func main() {
	var s []int
	for i:=0;i<100;i++{
		printArr(s)
		s= append(s,2*i+1)//自动扩容
	}
	fmt.Println(s)

	s1:=[]int{1,2,3,4,5}
	s2 :=make([]int ,16)//长度=容量
	s3:=make([]int,10,32)//长度 和 容量
	printArr(s1)
	printArr(s2)
	printArr(s3)
	fmt.Println("copy slice")
	copy(s2,s1)
	printArr(s2)
	fmt.Println("deleting slice")
	//s2[:3]+s2[4:] 删除下标为3的，需要重新append
	s2=append(s2[:3],s2[4:]...)
	printArr(s2)
//	删除头尾
	front:=s2[0]
	s2=s2[1:]
	fmt.Printf("pop from head ,head is %v\n,s2 is %v\n",front,s2)
	printArr(s2)
	tail := s2[len(s2)-1]
	s2=s2[:len(s2)-1]
	fmt.Printf("pop from tail ,tail is %v\n,s2 is %v\n",tail,s2)
	printArr(s2)




}
