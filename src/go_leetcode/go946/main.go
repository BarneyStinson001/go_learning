package main

import "fmt"

func validateStackSequences(pushed []int, popped []int) bool {
	l1:=len(pushed)
	l2:=len(popped)
	if l1!=l2{
		return false
	}
	if l1==1{
		return popped[0]==pushed[0]
	}

	//print(pushed[0])
	sta:=make([]int,1)
	sta[0]=pushed[0]


	i:=0
	for j:=0;j<l2&&i<l1;{
		fmt.Println("start",i,j,sta)
		l3:=len(sta)
		k:=0
		//i有问题  ,应该是len(sta)
		for ;l3-k>0 && j+k<l1 ;{
			if popped[j+k]==pushed[l3-1-k]{
				k+=1
			}else{
				break
			}
		}
		j+=k
		if j==l1{
			fmt.Println("end",i,j,sta)
			return true
		}
		if(l3==1 && k==1)
		sta=sta[0:l3-k]
		fmt.Println(i,j,sta)
		i+=1
		if  i>=l1{
			break
		}
		sta=append(sta,pushed[i])
		fmt.Println(i,j,sta)

	}
	// 0,1  1,0 else需要判断pushed[0] popped[-1]
	// 扩展   00001  10000
	if len(sta)==0{
		return true
	}else{
		for a:=0;a<len(sta);a++{
			if sta[a]==popped[l1-1-a]{
				continue
			}else{
				return false
			}
		}
		return true
	}

}

func main()  {
	a,b:=[]int{1,2,3,4,5},[]int{4,5,3,2,1}
	fmt.Println(validateStackSequences(a,b))
	c,d:=[]int{2,1,0},[]int{2,1,0}
	fmt.Println(validateStackSequences(c,d))

}