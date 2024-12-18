package main

import (
	"encoding/json"
	"fmt"
)

type Result struct {
	// Status 的首字母大写
	Status int
}

func main() {
	var data = []byte(`{"status":200}`)
	result := &Result{}
	if err := json.Unmarshal(data, &result); err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("result=%+v", *result)
}
