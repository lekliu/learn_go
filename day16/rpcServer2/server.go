package main

import (
	"errors"
	"log"
	"net"
	"net/rpc"
)

// 结构体，用于注册
type Arith struct{}

// 声明参数结构体
type ArithRequest struct {
	A, B int
}

// 返回客户端的结果
type ArithResponse struct {
	//乘积
	Pro int
	//商
	Quo int
	//余数
	Rem int
}

// 乘法
func (this *Arith) Multiply(req *ArithRequest, res *ArithResponse) error {
	res.Pro = req.A * req.B
	return nil
}

// 商和除法
func (this *Arith) Divide(req *ArithRequest, res *ArithResponse) error {
	if req.B == 0 {
		return errors.New("divide by zero")
	}
	res.Quo = req.A / req.B
	res.Rem = req.A % req.B
	return nil
}

// 主函数
func main() {
	rpc.Register(new(Arith))
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	//TODO
}
