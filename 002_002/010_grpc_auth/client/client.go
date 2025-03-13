package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	pb "studyProject/002_002/008_grpc_metadata/proto" // 替换为你的 proto 生成包
)

type CustomerAuth struct{}

func (ca *CustomerAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"authorization": "Wang token123"}, nil
}
func (ca *CustomerAuth) RequireTransportSecurity() bool {
	return false
}
func main() {

	opt := grpc.WithPerRPCCredentials(&CustomerAuth{})
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), opt)
	if err != nil {
		log.Fatalf("无法连接: %v", err)
	}
	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)

	// 创建 Metadata
	md := metadata.Pairs(
	// "authorization", "Wang token123", // 传递认证 Token
	// "trace-id", "abc123", // 传递 Trace ID
	)

	// 将 Metadata 附加到 context
	// 创建带 Metadata 的 Context
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	//
	var headerMd metadata.MD

	// 调用 gRPC 方法
	res, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Chen"}, grpc.Header(&headerMd))
	if err != nil {
		log.Fatalf("调用失败: %v", err)
	}

	fmt.Println("Response:", res.Message)

	// 获取服务端请求头 Metadata。
	// 读取 `server-version` 字段
	if serverVersion, exists := headerMd["server-version"]; exists {
		fmt.Println("server-version:", serverVersion)
	}
	// 读取 `response-time` 字段
	if responseTime, exists := headerMd["response-time"]; exists {
		fmt.Println("response-time:", responseTime)
	}
}
