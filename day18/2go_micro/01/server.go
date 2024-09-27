package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"log"
	pb "main/proto"
)

// 声明结构体
type Hello struct{}

func (this *Hello) Info(ctx context.Context, req *pb.InfoRequest, resp *pb.InfoResponse) (err error) {
	resp.Msg = "Hello " + req.Username
	return nil
}

func main() {
	// 1. 得到服务端实例
	service := micro.NewService(
		//设置微服务的名，用来访问
		//客户端访问： mico call hello  Hello.Info("username":"zhangsan")
		micro.Name("go.micro.service.hello"),
		micro.Version("latest"),
	)
	// 2. 初始化
	service.Init()
	// 3、 服务注册
	pb.RegisterHelloHandler(service.Server(), new(Hello))
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
	// 4.启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
