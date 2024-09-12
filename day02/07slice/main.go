package main

import "fmt"

func main() {
	var s1 []int
	a := [...]int{1, 2, 3, 4}
	s1 = a[1:4]
	fmt.Print(s1)
	a[3] = 100
	fmt.Print(s1)
}
