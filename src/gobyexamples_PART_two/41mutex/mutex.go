package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mu sync.Mutex
	counters map[string]int
}

func (c *Container)inc( name string)  {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	c:=Container{
		//mu:       sync.Mutex{},
		counters: map[string]int{"a":0,"b":0},
	}

	var wg sync.WaitGroup

	DoIncre := func(name string,n int) {
		for i:=0;i<n;i++{
			c.inc(name)
		}
		wg.Done()
	}
	wg.Add(4)
	go DoIncre("a",10000)
	go DoIncre("a",10000)
	go DoIncre("b",10000)
	go DoIncre("b",10000)

	wg.Wait()
	fmt.Println(c)
}
