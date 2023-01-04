package main

import (
	"fmt"
	"sync"
	"time"
)

var rw sync.RWMutex
var wg sync.WaitGroup
func read() {
	defer wg.Done()
	rw.RLock()
	fmt.Println("开始读取数据")
	time.Sleep(time.Second)
	fmt.Println("读取数据成功")
	rw.RUnlock()
}

func writer() {
	defer wg.Done()
	rw.Lock()
	fmt.Println("开始写入数据")
	time.Sleep(time.Second*10)
	fmt.Println("写入数据成功")
	rw.Unlock()
}

func main() {
	wg.Add(6)
	 for i:=0;i<5;i++{
	 	go read()
	 }

	 for j:=0;j<1;j++{
	 	go writer()
	 }
	 wg.Wait()
	 fmt.Println("end")

}