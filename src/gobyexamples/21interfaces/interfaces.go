package main

import (
	"fmt"
	"math"
)

//几何形状的接口：要包含面积和周长
//两个实现类
//都实现几何形中的两个方法
//接口变量有一个接口，就可以调用接口定义中的方法，不管是圆还是长方形
type geometry interface {
	area() float64
	len() float64
}

type circle struct {
	radius float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) len() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) len() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry)  {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.len())
}

func main() {
	r:=rect{
		width:  6,
		height: 8,
	}
	measure(r)
	c:=circle{radius: 10.0}
	measure(c)
}