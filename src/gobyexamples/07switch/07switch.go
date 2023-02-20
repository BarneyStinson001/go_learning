package main

import (
	"fmt"
	"time"
)

func main() {
	i:=2
	fmt.Print("write ",i," as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's weekend")
	default:
		fmt.Println("it's weekday")
		}

	t:=time.Now()
	switch  {
	case t.Hour() < 12:
		fmt.Println("before noon")
	default:
		fmt.Println("after noon")
	}

	WhatAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("it is int")
		case bool:
			fmt.Println("it is bool")

		default:
			fmt.Printf("the type is %T",t)
		}


	}
	WhatAmI(true)
	WhatAmI(1)
	WhatAmI("abc")

}

/*
case可以接多个值，用逗号隔开
默认：default分支
不加表达式，直接在case里判断  等价于ifslse
switch 不比较值，还可以比较类型，做类型判断：
 */
