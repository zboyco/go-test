package main

import (
	"context"
	"fmt"

	"github.com/zboyco/go-test/gRPC/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:9043", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)

	for i := 0; i < 10; i++ {
		res, err := c.Hello(context.TODO(),
			&proto.HelloReq{
				Name: fmt.Sprintf("baiyang-%v", i),
			},
		)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(res.Answer)
	}

}
