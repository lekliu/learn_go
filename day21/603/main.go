package main

import (
	"fmt"
	"time"
)

const zero = 0.0

func f1() {
	start := time.Now().UnixNano()
	for a := 0; a <= 500; a++ {
		for b := 0; b <= 500; b++ {
			c := 1000 - a - b
			if a*a+b*b == c*c {
				fmt.Printf("a:%d,b:%d,c:%d\n", a, b, c)
			}
		}
	}
	end := time.Now().UnixNano()
	fmt.Println(end - start)
}

//如果a+b+c=1000，且a×a+b×b=c×c （a,b,c为自然数)。
//如何求出所有a,b,c可能的组合?
//不使用数学公式和math包,代码实现,并且尽可能保证算法效率高效

func main() {
	f1()
	fmt.Printf("%T\n", zero)
}
