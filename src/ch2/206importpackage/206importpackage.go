package main

import (
	"fmt"
	"go_learning/src/ch2/205packageuse"
	"os"
	"strconv"
)

func main() {
	for _,arg := range os.Args[1:]{
		t,err :=strconv.ParseFloat(arg,64)
		if err!=nil{
			fmt.Fprintf(os.Stderr,"206importpackage:%v\n",err)
			os.Exit(1)
		}
		f:=tempconv.Fahrenheit(t)
		c:=tempconv.Celsius(t)
		fmt.Printf("%s = %s ,%s = %s \n",f,tempconv.FTOC(f),c,tempconv.CTOF(c))

	}
}


//go env
//set GOPATH=D:\github_learning\go_learning\go.mod
//src里放源码
//自定义包的导入：
//	设置变量GO111MODULE为on
//	go mod init + projectname：初始化mod
