package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// waitGroup

func f() {
	rand.Seed(time.Now().UnixNano()) //保证每次运行的随机数不同
	for i := 0; i < 5; i++ {
		r1 := rand.Int()
		r2 := rand.Intn(10) //0<= x <10
		fmt.Println(r1, r2)
	}
}

var wg sync.WaitGroup

func f1(i int) {
	defer wg.Done()
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(1300)))
	fmt.Println(i)
}

func main() {
	// f()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		f1(i)
	}
	wg.Wait() //等待wg的计数器降为0
}
