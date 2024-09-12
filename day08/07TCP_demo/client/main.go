package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// tcp client

func main() {
	// 1.与server端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("dial 127.0.0.1:20000 failed, err:", err)
		return
	}
	// 2. 发送数据
	var msg string
	// if len(os.Args) < 2 {
	// 	msg = "hello wangzi 55"
	// } else {
	// 	msg = os.Args[1]
	// }
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Please input:")
		msg, _ = reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg))
	}
	conn.Close()
}
