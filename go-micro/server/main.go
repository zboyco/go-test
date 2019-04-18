package main

import (
	"context"
	"fmt"

	"github.com/micro/go-micro"

	"github.com/zboyco/go-test/go-micro/proto"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloReq, res *proto.HelloRes) error {
	res.Result = "Hello " + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("greeter"),
	)

	service.Init()

	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
