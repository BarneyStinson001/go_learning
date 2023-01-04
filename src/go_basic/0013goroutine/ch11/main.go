package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

//var stop bool
var stop chan bool = make(chan bool, 1) //不能放在main函数中初始化make

//context方法
func cpuInfo(ctx context.Context) {
	defer wg.Done()
	for {
		//方案1
		//if stop{
		//	break
		//}

		//错误 ：  阻塞
		//if <-stop{
		//	break
		//}

		//方案2
		//select {
		//case <-stop:
		//	fmt.Println("退出CPU监控")
		//	return
		//default://也没default也会阻塞
		//	time.Sleep(time.Second * 1)
		//	fmt.Println("读取CPU信息")
		//}

		//方案3：
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

//需求0  启动协程监控CPU
//需求1: 协程循环监控CPU
//需求2：可以停掉监控CPU
//		方案1：全局变量 stop
//		方案2：通过chan
//更优雅的方法  方案3： context
func main() {

	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())

	go cpuInfo(ctx)

	time.Sleep(time.Second * 5)
	stop <- true
	cancel()
	wg.Wait()
	fmt.Println("信息监控完成")

}
