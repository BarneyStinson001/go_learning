package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	for k, v := range counts {
		if v > 1 {
			fmt.Printf("%d\t%s\n", v, k)
		}
	}
}
//make用来新建map
//input.scan读取下一行，去掉\n  input.Text()获取文本
//counts[input.text()]++  初始化默认值为0

//ctrl+z 结束标准输入


