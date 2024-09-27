package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc/proto"
	"net"
)

// 1. 需要监听
// 2. 需要实例化gRPC 服务端
// 3. 在gRPC上注册微服务
// 4. 启动服务器

// 定义空接口
type UserInfoService struct{}

var u = UserInfoService{}

// 实现方法
func (s *UserInfoService) GetUserInfo(ctx context.Context, req *pb.UserRequest) (resp *pb.UserResponse, err error) {
	// 通过用户名查询用户信息
	name := req.Name
	//从数据库中查用户信息
	if name == "zs" {
		resp = &pb.UserResponse{
			Id:    1,
			Name:  name,
			Age:   18,
			Hobby: []string{"Sing", "Run", "Others"},
		}
	}
	return
}

func main() {
	// 1地址
	addr := "127.0.0.1:8000"
	// 2监听
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("server listen err:", err)
	}
	fmt.Println("server listen addr:", addr)
	s := grpc.NewServer()
	// 3 在gRPC上注册微服务
	pb.RegisterUserInfoSercieServer(s, &u)
	// 4. 启动服务器
	s.Serve(lis)
}
