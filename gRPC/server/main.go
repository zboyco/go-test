package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"

	"github.com/zboyco/go-test/gRPC/proto"
)

type server struct{}

func (s *server) Hello(c context.Context, req *proto.HelloReq) (*proto.HelloRes, error) {
	return &proto.HelloRes{
		Answer: "Hello " + req.Name,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9043")
	if err != nil {
		fmt.Println(err)
		return
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})
	s.Serve(lis)
}
