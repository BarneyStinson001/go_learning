package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a+b
	case "-":
		return a-b
	case "*":
		return a*b
	case "/":
		return a/b
	default:
		panic("unsupported operation：" + op )
	}
}
func div(a, b int) (q int,r int) {
	q  = a/b
	r = a%b
	return
}

//switch比较土  ，换成函数式编程
func  apply(op func(int,int)int,a ,b int)  int{
	p :=reflect.ValueOf(op).Pointer()
	opname := runtime.FuncForPC(p).Name()
	fmt.Printf("call function %s with args "+"(%d,%d)",opname,a,b)
	return op(a,b )
}
func sum(numbers  ...int)	int	  {
	s:=0
	for i:=range numbers{
		s += numbers[i]
	}
	return s
}

func swap(a,b int){
	a,b=b,a
}

func swap_p(a ,b *int)  {
	*a ,*b=*b,*a
}
func swap_2(a,b int)(int,int){
	return b,a
}


func main() {
	fmt.Println(eval(1,2,"+"))
	fmt.Println(eval(1,2,"-"))
	fmt.Println(eval(1,2,"*"))
	fmt.Println(eval(1,2,"/"))
	//fmt.Println(eval(1,2,"**"))
	q,r := div(14,5)
	fmt.Println(q,r)
	fmt.Println(apply(func(a int, b int) int {
		return int(math.Pow(float64(a),float64(b)))
	},3,4))

	fmt.Println(sum(1,2,3,4))
	l,m := 3,4
	swap(l,m)
	fmt.Println(l,m)  //会失败


	o,p:=3,4
	swap_p(&o,&p)
	fmt.Println(o,p)

	r,s :=3,4
	r,s=swap_2(r,s)
	fmt.Println(r,s)
}
