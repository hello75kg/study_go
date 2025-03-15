package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	pb "studyProject/003_003/007_load_balance_grpc_consul/proto"
	"time"

	_ "github.com/mbobakov/grpc-consul-resolver"
)

func main() {
	// 连接 Consul 进行 gRPC 负载均衡
	conn, err := grpc.Dial(
		"consul://192.168.0.249:8500/grpc-service?wait=14s", // 连接 Consul
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`), // 轮询负载均衡
	)
	if err != nil {
		log.Fatalf("连接失败: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	for i := 0; i < 10; i++ {
		resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "wang"})
		if err != nil {
			log.Fatalf("调用失败: %v", err)
		}
		fmt.Println(resp.Message)
		time.Sleep(1 * time.Second) // 观察负载均衡效果
	}
}
