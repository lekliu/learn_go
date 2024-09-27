package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	c := make(chan int, 2)
	c <- 2
	data, err := json.Marshal(&c)
	fmt.Println(string(data), err)

	var i = 3
	b := &i
	data, err = json.Marshal(b)
	fmt.Println(string(data), err)
}
