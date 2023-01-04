package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	for i:=0;i<5;i++{
		wg.Add(2)
		go func(n int) {
			defer wg.Done()
			fmt.Println(n)
			//wg.Done()  //用defer
		}(i)
	}
	wg.Wait()//阻塞住
}
