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


func main() {
	//父context生成的子context，父cancel执行也会调用子cancel方法
	wg.Add(1)
	//ctx, _:= context.WithTimeout(context.Background(),time.Second*3)
	ctx, cancel:= context.WithTimeout(context.Background(),time.Second*3)

	go cpuInfo(ctx)
	time.Sleep(time.Second * 1)
	cancel()
	wg.Wait()
	fmt.Println("信息监控完成")

}
