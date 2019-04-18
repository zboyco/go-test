package main

import (
	"context"

	"github.com/micro/go-micro"

	"github.com/zboyco/go-test/go-micro/proto"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloReq, res *proto.HelloRes) error {
	res.Result = "Hello " + req.Name
	return nil
}

func main() {
	fuc := micro.NewFunction(
		micro.Name("greeter"),
	)

	fuc.Init()

	fuc.Handle(new(Greeter))

	fuc.Run()
}
