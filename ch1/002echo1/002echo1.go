package main

import (
	"fmt"
	"os"
)

func main() {
	var s,sep string
	for i:=1;i<len(os.Args);i++{
		s+=sep+os.Args[i]
		sep=" "
	}
	fmt.Println(s)
}

//os.Args
//for循环语句：
//for ini;condition;post{ }
//for condition{}   ==while 循环
//for {}  死循环 里面可以用break return终止