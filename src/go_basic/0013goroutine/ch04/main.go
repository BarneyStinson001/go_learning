package main

import (
	"fmt"
	"sync"
)

var total int
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		lock.Lock()
		total += 1
		lock.Unlock()
	}
}

func sub1() {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		lock.Lock()
		total -= 1
		lock.Unlock()
	}
}

func main() {
	wg.Add(2)
	go add()
	go sub1()
	wg.Wait()
	fmt.Println(total)
}
