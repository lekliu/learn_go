package design

import "fmt"

// 结构体
type Options struct {
	strOption1 string
	strOption2 string
	strOption3 string
	intOption1 int
	intOption2 int
	intOption3 int
}

// 初始化结构体
func InitOptions(strOption1 string, strOption2 string, strOption3 string,
	intOption1 int, intOption2 int, intOption3 int) {
	options := Options{}
	options.strOption1 = strOption1
	options.strOption2 = strOption2
	options.strOption3 = strOption3
	options.intOption1 = intOption1
	options.intOption2 = intOption2
	options.intOption3 = intOption3
	fmt.Printf("init Options: %v\n", options)
}

func main() {
	InitOptions("str1", "str2", "str3", 1, 2, 3)
}
