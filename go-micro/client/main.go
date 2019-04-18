package main

import (
	"context"
	"fmt"

	"github.com/micro/go-micro"
	"github.com/zboyco/go-test/go-micro/proto"
)

func main() {
	service := micro.NewService(micro.Name("greeter.client"))
	service.Init()

	greeter := proto.NewGreeterService("greeter", service.Client())

	res, err := greeter.Hello(context.TODO(),
		&proto.HelloReq{
			Name: "baiyang",
		},
	)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res.Result)
}
