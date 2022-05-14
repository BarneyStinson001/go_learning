package main
//输出os.Args[0],即命令的命令
//输出参数的索引和值，每行一个
//尝试测试可能低效的程序和使用strings.Join的程序在执行时间上的差异  1.6节time包  11.4节写性能评估测试

import "fmt"
import "os"

func main() {
	fmt.Println(os.Args[0])
	for i:=1;i<len(os.Args);i++ {
		fmt.Println(i, os.Args[i])
	}

}
