package main

import "fmt"

func main() {
	var a int = 10
	fmt.Println(a)
	{
		a := 9
		fmt.Println(a)
		a = 8
	}
	fmt.Println(a)
}

//A. 10 10 10
//B. 10  9  9
//C. 10  9 10
//D. 10  9  8
