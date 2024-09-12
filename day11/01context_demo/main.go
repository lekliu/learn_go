package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var notify bool

//为什么需要context

func f() {
	defer wg.Done()
	for {
		fmt.Println("hehe")
		time.Sleep(time.Millisecond * 500)
		if notify {
			break
		}
	}
}

func main() {
	wg.Add(1)
	go f()
	time.Sleep(time.Second * 5)
	notify = true
	wg.Wait()
	// 如何通知子goroutine退出
}
