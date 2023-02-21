package main

import "fmt"

/*
结构体嵌套
嵌套的结构体的字段可以直接在外面访问
结构体里的结构体实现的方法，可以被直接调用

结构体里的结构体实现的接口，外结构体也算实现了该接口

*/

type base struct {
	num int
}

func (b *base) decr() string {
	return fmt.Sprintf("base num is %v:", b.num)

}

type container struct {
	base
	str string
}

func main() {
	c := container{
		base: base{num: 3,},
		str:  "zhx",
	}
	fmt.Printf("co={num: %v,str: %v}", c.num, c.str)

	fmt.Println("also num:", c.base.num)

	fmt.Println(c.decr())

	type decr12 interface {
		decr() string
	}

	var d decr12 = &c//注意传的是类型还是指针类型
	fmt.Println("decribe: ", d.decr())

}
