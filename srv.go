package main

import (
	"context"
	"fmt"

	micro "github.com/micro/go-micro/v2"
	proto "greet/proto" //这里写你的proto文件放置路劲
)

type Greeter struct{} //此处的Greeter必须跟proto文件里的service 一致，并且需要实现protoc文件里的rpc接口

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("greeter"), //此处的名字随意，客户端调用时需要跟此处的一支
	)

	// Init will parse the command line flags.
	service.Init()

	// Register handler
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
