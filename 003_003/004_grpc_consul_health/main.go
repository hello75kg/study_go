package main

import (
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
	"time"
)

func Register(address string, port int, name string, tags []string) error {
	cfg := api.DefaultConfig()
	cfg.Address = "192.168.0.249:8500"

	client, _ := api.NewClient(cfg)

	registration := new(api.AgentServiceRegistration)
	registration.ID = address
	registration.Name = name
	registration.Tags = tags
	registration.Port = port
	check := api.AgentServiceCheck{
		GRPC:                           "192.168.0.249:50051",
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "5s",
	}
	registration.Check = &check

	err := client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
	return nil
}

func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}

	go func() {
		grpcServer := grpc.NewServer()
		grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())

		log.Println("gRPC 服务器运行在端口 50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("启动 gRPC 服务器失败: %v", err)
		}
	}()

	_ = Register("192.168.0.249", 50051, "user-web", []string{"wshop-web", "wang"})

	time.Sleep(10 * time.Second)
}
