package main

import (
	"net"
	"fmt"
	"context"

	"demo/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":7899")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	c := &center{}
	pb.RegisterCenterServer(s, c)
	reflection.Register(s)
	s.Serve(lis)
}

type center struct {
	pb.UnimplementedCenterServer
}

func (c *center) Call(ctx context.Context, req *pb.CallReq) (*pb.CallResp, error) {
	fmt.Println("receive call", req)
	return &pb.CallResp{Ok: true}, nil
}


