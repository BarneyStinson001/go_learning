package main

import (
	"flag"
	"log"
)

var name string
func main()  {
	flag.Parse()
	goCmd := flag.NewFlagSet("go",flag.ExitOnError)
	//flag.StringVar(&name,"name","lisi","指定名字")
	goCmd.StringVar(&name,"name","lisi","指定名字")

	phpCmd :=flag.NewFlagSet("php",flag.ExitOnError)
	phpCmd.StringVar(&name,"n","lisi","指定名字")

	args :=flag.Args()
	switch args[0]	 {
	case "go":
		_=goCmd.Parse(args[1:])
	case "php":
		_=phpCmd.Parse(args[1:])
	}
	log.Printf("name: %s",name)
}

//D:\github_learning\go_learning\src\go-programming-tour-book\tour\ch1\ch1.1.3.2>go run ch1_1_1_3_2.go go -name zhx
//2022/06/08 00:05:40 name: zhx
//
//D:\github_learning\go_learning\src\go-programming-tour-book\tour\ch1\ch1.1.3.2>go run ch1_1_1_3_2.go php -n zhx
//2022/06/08 00:05:51 name: zhx
//
//D:\github_learning\go_learning\src\go-programming-tour-book\tour\ch1\ch1.1.3.2>go run ch1_1_1_3_2.go go -n zhx
//flag provided but not defined: -n
//Usage of go:
//  -name string
//        指定名字 (default "lisi")
//exit status 2