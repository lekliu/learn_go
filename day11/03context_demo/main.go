package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var exitChan chan bool = make(chan bool, 1)

// 为什么需要context
func f2(ctx context.Context) {
	defer wg.Done()
FORLOOP:
	for {
		fmt.Println("hehe22")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done(): //等待上级通知
			break FORLOOP
		default:
		}
	}
}

func f(ctx context.Context) {
	defer wg.Done()
	wg.Add(1)
	go f2(ctx)
FORLOOP:
	for {
		fmt.Println("hehe")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done(): //等待上级通知
			break FORLOOP
		default:
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go f(ctx)
	time.Sleep(time.Second * 5)
	cancel() //通知子Goroute 退出
	wg.Wait()
	// 如何通知子goroutine退出
}
