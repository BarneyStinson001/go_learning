package main

import "fmt"

func main() {

	m:=make(map[string]int)

	m["zhangsan"]=80
	m["lisi"]=85

	fmt.Println("map:",m)

	v:=m["lisi"]
	fmt.Println("v: ",v)

	fmt.Println(len(m))

	delete(m,"lisi")
	fmt.Println("delete lisi",len(m))

	if va,ok:=m["lisi"];ok{
		print(va)
	}else {
		fmt.Println(ok,m)
	}


	m1:=map[string]int{"wangwu":90}
	fmt.Println(m1)
}
