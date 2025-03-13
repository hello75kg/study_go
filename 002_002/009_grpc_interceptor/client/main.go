package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"studyProject/002_001/005_grpc/proto"
	"time"
)

func main() {
	interceptor := func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		timeStart := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		fmt.Println(time.Since(timeStart))
		return err
	}
	opt := grpc.WithUnaryInterceptor(interceptor)
	conn, _ := grpc.NewClient("127.0.0.1:8088", grpc.WithTransportCredentials(insecure.NewCredentials()), opt)
	defer conn.Close()
	client := proto.NewGreeterClient(conn)
	res, _ := client.SayHello(context.Background(), &proto.HelloRequest{
		Name: "grpc",
	})
	fmt.Println(res)

}
