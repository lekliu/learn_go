package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	//开启3个goroutine
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}
	// 5个任务
	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs)

	//输出结果
	for a := 1; a <= 5; a++ {
		<-results
	}
}
