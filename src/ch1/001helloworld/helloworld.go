package main

import "fmt"

func main()  {
	fmt.Println("hello, world")

}
//go语言通过包进行管理，首先指明这个文件属于哪个包
//main包特殊之处，用来定义一个可独立执行的文件，而不是库
//main函数则是程序最开始的地方


//import部分代表需要导入的包，编译器源文件需要哪些包
//不能导入没有使用的包，会编译失败。


