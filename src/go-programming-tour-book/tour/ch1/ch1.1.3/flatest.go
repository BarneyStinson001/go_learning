package main

import (
	"flag"
	"log"
)

func main() {
	var name,age string
	flag.StringVar(&name,"name","lisi","指定名字")
	flag.StringVar(&age,"age","18","指定年龄")
	flag.Parse()

	log.Printf("name: %s  age: %s",name,age)
}

//D:\github_learning\go_learning\src\go-programming-tour-book\tour\ch1\ch1.1.3>go run flatest.go --name=zhx -age=19
//2022/06/07 23:45:46 name: zhx  age: 19
//
//D:\github_learning\go_learning\src\go-programming-tour-book\tour\ch1\ch1.1.3>go run flatest.go --name=zhx --age=19
//2022/06/07 23:45:50 name: zhx  age: 19
//
//D:\github_learning\go_learning\src\go-programming-tour-book\tour\ch1\ch1.1.3>go run flatest.go -name=zhx --age=19
//2022/06/07 23:45:55 name: zhx  age: 19
//
//D:\github_learning\go_learning\src\go-programming-tour-book\tour\ch1\ch1.1.3>go run flatest.go -name=zhx -age=19
//2022/06/07 23:46:02 name: zhx  age: 19