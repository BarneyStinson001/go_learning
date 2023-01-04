package main

import "fmt"

func updateSlice(s []int)  {
	s[0]=200
}

func main() {
	arr := [...]int{1,2,3,4,5,6,7,8,}
	fmt.Println("arr[2:6] =",arr[2:6])
	fmt.Println("arr[:6] =",arr[:6])
	fmt.Println("arr[2:] =",arr[2:])
	fmt.Println("arr[:] =",arr[:])

	s1 := arr[2:6]
	updateSlice(s1)
	fmt.Println("s =",s1)
	fmt.Println("arr = ",arr)

	s2:= arr[:]
	updateSlice(s2)
	fmt.Println("s =",s2)
	fmt.Println("arr = ",arr)

	s3:=arr[:5]
	fmt.Println(s3)
	s3=s3[2:]
	fmt.Println(s3)

	s4:=arr[2:6]
	fmt.Printf("s4 = %v ,len = %d,cap = %d\n",s4,len(s4),cap(s4))
	s5:=s4[3:5]
	fmt.Printf("s5 = %v ,len = %d,cap = %d\n",s5,len(s5),cap(s5))
}
