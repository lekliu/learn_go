package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type job struct {
	value int64
}

type result struct {
	job *job
	sum int64
}

var wg sync.WaitGroup

func zhoulin(zl chan<- *job) {
	defer wg.Done()
	for {
		intRand := rand.Int63()
		zl <- &job{
			value: intRand,
		}
		time.Sleep(time.Millisecond * 500)
	}
}

func baodelu(zl <-chan *job, resultChan chan<- *result) {
	defer wg.Done()
	// 从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
	for {
		job := <-zl
		sum := int64(0)
		n := job.value
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		newResult := &result{
			job: job,
			sum: sum,
		}
		resultChan <- newResult
	}
}

func main() {
	var jobChan = make(chan *job, 50)
	var resultChan = make(chan *result, 50)

	wg.Add(1)
	go zhoulin(jobChan)
	wg.Add(24)
	for i := 0; i < 24; i++ {
		go baodelu(jobChan, resultChan)
	}

	for {
		result := <-resultChan
		fmt.Println(result.job.value, result.sum)
	}
}
