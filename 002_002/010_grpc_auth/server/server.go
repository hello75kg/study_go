package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	pb "studyProject/002_002/008_grpc_metadata/proto"
)

// 实现 gRPC 服务
type HelloService struct {
	pb.UnimplementedHelloServiceServer
}

func (s *HelloService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	// 获取客户端发送的 Metadata。
	// 解析 Metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("未找到 Metadata")
	} else {
		// 读取 `authorization` 字段
		if auth, exists := md["authorization"]; exists {
			fmt.Println("Authorization:", auth)
		}
		// 读取 `trace-id` 字段
		if traceID, exists := md["trace-id"]; exists {
			fmt.Println("Trace ID:", traceID)
		}
	}

	md2 := metadata.Pairs(
		"server-version", "1.0.0",
		"response-time", "100ms",
	)
	_ = grpc.SendHeader(ctx, md2) // 发送响应头部 Metadata
	return &pb.HelloResponse{Message: "Hello, " + req.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}
	res := func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return resp, status.Error(codes.Unauthenticated, "无认证信息")
		}

		var token string
		if res, ok := md["authorization"]; ok {
			token = res[0]
		}
		if token != "Wang token123" {
			return resp, status.Error(codes.Unauthenticated, "验证失败")
		}

		res, err := handler(ctx, req)

		return res, err
	}
	opt := grpc.UnaryInterceptor(res)
	grpcServer := grpc.NewServer(opt)
	pb.RegisterHelloServiceServer(grpcServer, &HelloService{})

	log.Println("gRPC 服务器运行在端口 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("启动 gRPC 服务器失败: %v", err)
	}
}
