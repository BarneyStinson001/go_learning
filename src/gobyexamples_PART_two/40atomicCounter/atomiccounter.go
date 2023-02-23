package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var ops uint64

	var wg sync.WaitGroup

	for i:=0;i<50;i++{
		wg.Add(1)
		go func() {
			defer wg.Done()
			for c := 0;c<1000;c++ {
				//使用`atomic.LoadUint64`. 这样的函数，当他们被更新时有读取也是安全的。
				atomic.AddUint64(&ops,1)
			//	如果不是用非原子操作，就会有数据竞争风险
			//used the non-atomic ，	data race failures
			}
		}()
	}

	wg.Wait()

	fmt.Println("ops:",ops)
}
