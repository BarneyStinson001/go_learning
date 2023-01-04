package main

import (
	"fmt"
	"time"
)

func main() {
	now:=time.Now()
	//time转化为字符串
	fmt.Println(now.Format("2006-01-02-15-04-05"))
	//字符串转time
	t1,_:=time.ParseInLocation("20060102150405",now.Format("20060102150405"),time.Local)
	fmt.Println(t1)

	//duration
	t,_:=time.ParseDuration("-1m")
	fmt.Println(t1.Add(t))
	fmt.Println(t1)

	//since
	time.Sleep(2)
	fmt.Println(time.Since(t1))

	//日期比较前后  是之前还是之后


	//日期加减duration
	

}
