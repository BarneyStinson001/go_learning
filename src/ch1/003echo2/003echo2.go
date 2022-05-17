package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

//go不允许蔡遵无用的临时变量。所以用_空标识符代替
//变量定义通常使用前两种
//s := ""
//var s string
