package main

import (
	"bufio"
	"fmt"
	"os"
)

//从文件列表读取，找出重复的行
func main()  {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files)==0{
		countLine(os.Stdin,counts)
	}else {
		for _,file :=range files{
			f,err :=os.Open(file)
			if err !=nil {
				fmt.Fprintf(os.Stderr,"dup2: %v\n",err)
				continue
			}
			countLine(f,counts)
			f.Close()
		}
	}
	fmt.Printf("n\tstr\n")
	for line,n :=range counts{
		if n>1{
			fmt.Printf("%d\t%s\n",n,line)
		}
	}
}
func countLine(file *os.File,count map[string]int)  {
	input :=bufio.NewScanner(file)
	for input.Scan() {
		count[input.Text()]++
	}
}
//go run 006dup2.go  从标准输入
//go run 006dup2.go 1.txt 2.txt 3.txt 4.txt
