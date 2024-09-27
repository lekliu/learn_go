package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"
	"log"
	pb "main/proto"
)

type Example struct{}
type Foo struct{}

func (e *Example) call(ctx context.Context, req *pb.CallRequest, resp *pb.CallResponse) (err error) {
	log.Print("收到Example请求")
	if len(req.Name) == 0 {
		return errors.BadRequest("go.micro.api,example", "no name")
	}
	resp.Msg = "Example.call 接收到了你的请求  " + req.Name
	return nil
}

func (f *Foo) Bar(ctx context.Context, req *pb.EmptyRequest, resp *pb.EmptyResponse) (err error) {
	log.Print("收到Foo.Bar请求")
	return nil
}

func main() {
	// 1. 得到服务端实例
	service := micro.NewService(
		//设置微服务的名，用来访问
		//客户端访问： mico call hello  Hello.Info("username":"zhangsan")
		micro.Name("go.micro.api.example"),
	)
	// 2. 初始化
	service.Init()
	// 3、 服务注册
	pb.RegisterExampleHandler(service.Server(), new(Example))
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
	pb.RegisterFooHandler(service.Server(), new(Foo))
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
	// 4.启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
