package main

import (
	"bufio"
	"fmt"
	"os"
)

// 打开文件写入内容
func writeDemo1() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("File Open failed,err:%v", err)
		return
	}
	//write
	fileObj.Write([]byte("hello bie"))
	fileObj.WriteString("\nstart coding")
	fileObj.Close()
}

func writeDemo2() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("File Open failed,err:%v", err)
		return
	}
	//创建一个写对象
	wr := bufio.NewWriter(fileObj)
	wr.WriteString("bufio start write data \n")
	wr.Flush()
}

func writeDemo3() {
	str := "This data is prepared for writing file."
	err := os.WriteFile("./xx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Printf("File Open failed,err:%v", err)
		return
	}

}

func main() {
	// writeDemo1()
	// writeDemo2()
	writeDemo3()
}
