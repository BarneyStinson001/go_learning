package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[1:])
	fmt.Printf("hello,world")
	os.Exit(-1)
}