package main

import (
	"context"
	"google.golang.org/grpc"
	"net"
	"studyProject/002_001/005_grpc/proto"
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
	g := grpc.NewServer()
	proto.RegisterGreeterServer(g, &Server{})
	listen, _ := net.Listen("tcp", ":8088")
	_ = g.Serve(listen)

}
