package main

import (
	"fmt"
)

func main() {
	s:=make([]string,3)
	fmt.Println("empty: " ,s)


	s[0]="avs"
	s[1]="scsdca"
	s[2]="sdasfcvcdv"
	fmt.Println("set: ",s)
	fmt.Println("get :",s[2])

	fmt.Println("len: ",len(s))

	s=append(s,"4")
	s=append(s,"5","6")
	fmt.Println("append: ",s)


	c:=make([]string,len(s))
	copy(c,s)

	fmt.Println("copy: ",c)

	l1:=s[2:4]
	l2:=s[:3]
	l3:=s[3:]
	fmt.Println(l1,l2,l3)


	t:=[]string{"1","2","3,"}
	fmt.Println(t)

	twoD:=make([][]int,3)
	for i:=0;i<3;i++{
		for j:=0;j<i;j++{
			twoD[i][j]=i+j
		}
	}
	fmt.Println("2d: ",twoD)
}

/*
切片比数组会更灵活
使用make进行创建
append  和扩容

复制


 */