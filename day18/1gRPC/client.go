package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc/proto"
)

// 1.连接服务端
// 2.实例化gRPC客户端
// 3.调用

func main() {
	// 1.连接服务端
	conn, err := grpc.Dial("127.0.0.1:8000", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("fail to dial: %v\n", err)
	}
	defer conn.Close()
	// 2.实例化gRPC客户端
	client := pb.NewUserInfoSercieClient(conn)
	// 3.组装请求参数，调用
	req := new(pb.UserRequest)
	req.Name = "zs"
	resp, err := client.GetUserInfo(context.Background(), req)
	if err != nil {
		fmt.Printf("fail to call GetUserInfo: %v\n", err)
	}
	fmt.Printf("响应结果: %+v\n", resp)
}
