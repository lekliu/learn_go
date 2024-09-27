package main

import "fmt"

type S struct {
	name string
}

func main() {
	// map 里结构体无法直接寻址，必须取址
	// 出错写法： m := map[string]S{"x":S{"one"}}
	m := map[string]*S{"x": &S{"one"}}
	m["x"].name = "two"
	fmt.Println(m["x"].name)
}
