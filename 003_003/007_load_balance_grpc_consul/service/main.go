package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	pb "studyProject/003_003/007_load_balance_grpc_consul/proto"
)

type server struct {
	pb.UnimplementedGreeterServer
	port int
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: fmt.Sprintf("Hello %s from port %d!", req.Name, s.port)}, nil
}

// 注册 gRPC 服务到 Consul
func registerServiceToConsul(port int) {
	config := api.DefaultConfig()
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("创建 Consul 客户端失败: %v", err)
	}

	reg := &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("grpc-server-%d", port),
		Name:    "grpc-service",
		Address: "192.168.0.249",
		Port:    port,
		Check: &api.AgentServiceCheck{
			GRPC:                           fmt.Sprintf("192.168.0.249:%d", port),
			Interval:                       "10s",
			Timeout:                        "5s",
			DeregisterCriticalServiceAfter: "30s",
		},
	}

	err = client.Agent().ServiceRegister(reg)
	if err != nil {
		log.Fatalf("注册服务到 Consul 失败: %v", err)
	}
	log.Printf("服务注册成功: grpc-service @ 192.168.0.249:%d", port)
}

func main() {
	port := 50054 // 运行多个实例时，改成不同端口
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}

	// 注册服务到 Consul
	registerServiceToConsul(port)

	s := grpc.NewServer()
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())
	pb.RegisterGreeterServer(s, &server{port: port})

	log.Printf("gRPC 服务器在端口 %d 启动", port)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}
