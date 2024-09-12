package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// UDP 客户端
func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("连接服务端失败，err:", err)
		return
	}
	defer socket.Close()
	// sendData := []byte("Hello server")
	// _, err = socket.Write(sendData) // 发送数据
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("请输入内容：")
		msg, _ := reader.ReadString('\n')
		_, err = socket.Write([]byte(msg)) // 发送数据
		if err != nil {
			fmt.Println("发送数据失败，err:", err)
			return
		}
		// 收回复的数据
		data := make([]byte, 4096)
		n, remoteAddr, err := socket.ReadFromUDP(data) // 接收数据
		if err != nil {
			fmt.Println("接收数据失败，err:", err)
			return
		}
		fmt.Printf("recv:%v addr:%v count:%v\n", string(data[:n]), remoteAddr, n)

	}
}
