package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func variable()  {
	var a  int
	var s  string
	var b = 1
	var	ss = "ss"
	c := 22
	fmt.Printf("%d %s \n" ,a, s)
	fmt.Println(b,ss,c)
}

func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))

	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi)+1)
	//欧拉公式 e的i*pi次方+1=0

}
func triangle() {
	d,e := 3,4
	var f int
	f = int(math.Sqrt(float64(d*d+e*e)))
	fmt.Println(f)
}
func consts() {
	const name = "abs"
	const age =  12
	fmt.Println(name,age)
}

func enums() {
	const (
		cpp = iota
		_
		python
		golang
	)
	const (
		b = 1<<(10*iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp,python,golang)
	fmt.Println(b,kb,mb,gb,tb,pb)
}

func main() {
	fmt.Println("Hello")
	variable()

	euler()

	triangle()

	consts()

	enums()
}

