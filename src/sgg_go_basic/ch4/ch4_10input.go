package main


import (
	"fmt"
)

func main()  {
	var name string
	var age int
	var sex	bool
	var addr string

	fmt.Println("input your name:\t")
	fmt.Scanln(&name)

	fmt.Println("input your age:\t")
	fmt.Scanln(&age)

	fmt.Println("input your gender:(true for male,false for female)")
	fmt.Scanln(&sex)

	fmt.Println("input your addr:\t")
	fmt.Scanln(&addr)

	fmt.Printf("name is %v\nage is %v\nismale: %v\naddr is %v\n",name,age,sex,addr)

	fmt.Println("please input your name,age,ismale,addr:(e.g. lisi,21,true,china)")
	fmt.Scanf("%s %d %t %s",&name,&age,&sex,&addr)
	fmt.Printf("name is %v\nage is %v\nismale: %v\naddr is %v\n",name,age,sex,addr)

}