package main

import "fmt"

func main() {
	var i int = 7
	b := &i
	fmt.Println(*b)
	b = b + 1
	b++
	b--
	fmt.Println(*b)
}
