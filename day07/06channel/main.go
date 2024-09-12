package main

import (
	"fmt"
	"sync"
)

// channel

var b chan int
var wg sync.WaitGroup

func noBuffChannel() {
	b = make(chan int) //不带缓冲区通道的初始化
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b
		fmt.Println("后台Goroutine从通道b接收数据 ", x)
	}()
	b <- 10
	fmt.Println("10发送到通道b中了....")
	wg.Wait()
}

func buffChannel() {
	b = make(chan int, 16) //带缓冲的通道的初始化
	b <- 10
	fmt.Println("10发送到通道b中了....")
	x := <-b
	fmt.Println("从通道b接收数据 ", x)
	close(b)
}

func main() {
	fmt.Println(b) //nil
	buffChannel()
}
