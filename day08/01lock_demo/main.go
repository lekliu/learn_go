package main

import (
	"fmt"
	"sync"
)

//锁

var x = 0
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	for i := 0; i < 50000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}

func main() {
	wg.Add(3)
	go add()
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
