package main

import (
	"log"
	"net/http"
	"sync"
	"time"
)

var logger = log.Default()

func mpdtest(wg *sync.WaitGroup, url string) int {
	defer wg.Done()
	rsp, err := http.Get(url)
	if err != nil {
		logger.Printf("get %s  failed:%s", url, err.Error())
	}
	logger.Printf("get %s statuscode:%d,contentLength:%d",url,rsp.StatusCode,rsp.ContentLength)
	defer rsp.Body.Close()
	return int(rsp.ContentLength)
}

func main() {
	//wg防止主进程退出
	urls := [10][100]string{}
	for i := range urls {
		for j := range urls[i] {
			urls[i][j] = "https://www.baidu.com/"
		}
	}

	var wg sync.WaitGroup
	for r := 0; r < 100; r++ {
		wg.Add(1000)
		for i := 0; i < 10; i++ {
			now := time.Now()
			for j := 0; j < 100; j++ {
				go mpdtest(&wg, urls[i][j]) //协程并发
			}
			log.Printf("time spent %v", time.Since(now))
			time.Sleep(time.Second * 3)
		}
		wg.Wait()
	}
}