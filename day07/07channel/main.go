package main

import (
	"fmt"
	"sync"
)

// channel练习
// 1. 启动一个goroutine，生成100个数发送到ch1
// 2. 启动一个goroutine,从ch1中取值，计算其平方放到ch2中
// 3. 在main中 从ch2取值打印出来

func f1(ch1 chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

func f2(ch1 chan int, ch2 chan int) {
	for {
		n, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- n * n
	}
	close(ch2) // 关闭后，channel不能写入，只能读取，读到结束返回nil
}

func main() {
	var ch1 = make(chan int, 50)
	var ch2 = make(chan int, 50)
	var wg sync.WaitGroup
	wg.Add(1)
	go f1(ch1, &wg)
	go f2(ch1, ch2)

	for {
		m, ok := <-ch2
		if !ok {
			break
		}
		fmt.Println(m)
	}
	wg.Wait() //本未例中可以不需要wg
}
