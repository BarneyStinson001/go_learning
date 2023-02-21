package main

import "fmt"

/*
给struct定义方法
 方法名前面是接收者 ，可以用类型也可以是指针

go在方法调用时自动转换指针和数据
用指针类型可以避免值传递拷贝引起的开销，以及允许方法修改接收者的结构

 */

type rect struct {
	width,height int
}

func (r *rect)area()  int{
	return r.width*r.height
}

func (r rect)leng() int	 {
	return 2*(r.height+r.width)
}

func main() {
	r:=rect{
		width:  6,
		height: 8,
	}
	fmt.Println("area: ",r.area())
	fmt.Println("len: ",r.leng())

	rp:=&r
	fmt.Println("area: ",rp.area())
	fmt.Println("len: ",rp.leng())

}
