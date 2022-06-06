package main

import (
	"flag"
	"fmt"
	"strings"
)

var n=flag.Bool("n",false,"omit trailing newline")
var sep=flag.String("s"," ","separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

//编译后执行

//.\203echo4.exe  -a   = -h  -help

//.\203echo4.exe  -s / a bc def  =  a/bc/def
//.\203echo4.exe  -n a  bc def  = a bc def 不换行
//.\203echo4.exe  a bc d       =  a bc def 换行·
