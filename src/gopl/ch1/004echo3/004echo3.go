package main

import (
	"fmt"
"os"
"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:]," "))
	fmt.Println(os.Args[1:])
}

//老方法，旧数据会生成垃圾，被垃圾回收，影响性能。
//面对大量数据时，代价太大。

//切片输出带[]
