package main

import (
	"errors"
	"fmt"
)


func f1(arg int)(int,error)  {
	//error是一个内置的接口，只要实现Error()方法就可以实现了这个接口
	if arg==19{
		return -1,errors.New("cant't work witg 19")//用errors.New来构造一个消息
	}
	return arg+1,nil//return nil,表示没有错误
}


type argError struct {
	arg int
	prob string
}
//argError为自定义类型，实现Error()方法就可以
func (e *argError)Error() string  {
	return fmt.Sprintf("%d - %s",e.arg,e.prob )
}

func f2(arg int)(int,error)  {
	if arg==19{
		return -1,&argError{
			arg:  arg,
			prob: "can't work with it",
		}
	}
	return arg+1,nil
}

func main()  {
	for _,i:=range []int{15,19}{
		if r,e:=f1(i);e!=nil{
			//if语句可以用表达式，后面直接跟条件判断，变量的作用域只在if作用域
			fmt.Println("f1 failed: ",e)
		}else {
			fmt.Println("f1 worked: ", r)
		}
	}
	for _,i:=range []int{15,19}{
		if r,e:=f2(i);e!=nil{
			fmt.Println("f2 failed:",e)
		}else {
			fmt.Println("f2 worked:", r)
		}
	}
	_,e:=f2(19)
	if ae,ok:=e.(*argError);ok{
		//类型断言，获取实例
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}