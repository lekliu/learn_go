package main

import (
	"fmt"
	"net"
)

// tcp server端
func processConn(conn net.Conn) {
	defer conn.Close()
	// 3、与客户端通信
	var tmp [128]byte
	for {
		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Println("read from conn failed, err:", err)
			return
		}
		fmt.Println(string(tmp[:n]))
	}
}

func main() {
	// 1、本地端口启动服务
	listerner, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("start tcp server on 127.0.0.1:20000 failed, err:", err)
		return
	}
	// 2、等待别人来跟我连接
	for {
		conn, err := listerner.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			return
		}

		go processConn(conn)

	}
}
