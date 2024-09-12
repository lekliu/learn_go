package main

import (
	"fmt"
	"sync"
	"time"
)

// rwlock
var (
	x      = 0
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwLock sync.RWMutex
)

func read() {
	defer wg.Done()
	// lock.Lock()
	rwLock.RLock()
	fmt.Println(x)
	time.Sleep(time.Millisecond)
	// lock.Unlock()
	rwLock.RUnlock()
}

func write() {
	defer wg.Done()
	// lock.Lock()
	rwLock.Lock()
	x = x + 1
	time.Sleep(time.Millisecond * 5)
	// lock.Unlock()
	rwLock.Unlock()
}

func main() {
	start := time.Now()

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go write()

	}

	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go read()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
