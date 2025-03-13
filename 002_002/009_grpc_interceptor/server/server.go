package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"studyProject/002_001/005_grpc/proto"
	"time"
)

type Server struct {
	proto.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{
		Message: "hello" + req.Name,
	}, nil
}

func main() {
	interceptor := func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		fmt.Println("server interceptor start")
		timeStart := time.Now()
		res, _ := handler(ctx, req)
		fmt.Println("server interceptor end", time.Since(timeStart))
		return res, err
	}
	opts := grpc.UnaryInterceptor(interceptor)
	g := grpc.NewServer(opts)
	proto.RegisterGreeterServer(g, &Server{})
	listen, _ := net.Listen("tcp", ":8088")
	_ = g.Serve(listen)

}
