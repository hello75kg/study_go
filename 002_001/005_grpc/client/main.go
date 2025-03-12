package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"studyProject/002_001/005_grpc/proto"
)

func main() {
	conn, _ := grpc.NewClient("127.0.0.1:8088", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	client := proto.NewGreeterClient(conn)
	res, _ := client.SayHello(context.Background(), &proto.HelloRequest{
		Name: "grpc",
	})
	fmt.Println(res)

}
