package main

import "fmt"

/*
struct 是一组fields的集合，有效地把数据组合在一起
 */

type Person struct {
	name string
	age int
}

func NewPerson(name string)*Person  {
	p:=Person{
		name: name,
		age:  0,
	}
	p.age=20
	return &p
}

func main() {
	fmt.Println(Person{"zhangsan",20})
	//指定字段进行初始化
	fmt.Println(Person{
		name: "lisi",
		age:  20,
	})
	//省略的字段初始化为零值
	fmt.Println(Person{name: "wangwu"})

	fmt.Println(&Person{name: "liliu",age: 25})
//	习惯用法 封装NewXxxx函数来构造结构体
	fmt.Println(NewPerson("zhuba"))

	p:=Person{
		name: "abdc",
		age:  25,
	}
	fmt.Println(p.age)

	pp:=&p
	fmt.Println(pp.name)

	pp.age=50
	fmt.Println(p.age)

}