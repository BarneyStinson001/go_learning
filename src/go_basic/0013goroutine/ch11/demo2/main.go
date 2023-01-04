package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)
var wg sync.WaitGroup

func cpuInfo(ctx context.Context) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("退出CPU监控")
			return
		default:
			time.Sleep(time.Second * 1)
			fmt.Println("读取CPU信息")
		}
	}
}

func memInfo(ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("退出内存监控")
			return
		default:
			time.Sleep(time.Second * 4)
			fmt.Println("读取内存信息")
		}
	}
}

func main() {
//父context生成的子context，父cancel执行也会调用子cancel方法
	wg.Add(2)
	ctx, cancel := context.WithCancel(context.Background())
	go cpuInfo(ctx)
	go memInfo(ctx)//这个放在cpuInfo里执行，也是可以的只要传递ctx
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
	fmt.Println("信息监控完成")

}
