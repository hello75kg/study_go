package main

import (
	"context"
	"io"
	"log"
	"strconv"
	"time"

	"google.golang.org/grpc"
	pb "studyProject/002_001/006_grpc_stream/proto"
)

func main() {
	conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())

	defer conn.Close()
	client := pb.NewStreamServiceClient(conn)

	// 1. 客户端流调用
	clientStreamUpload(client)

	// 2. 服务端流调用
	serverStreamDownload(client)

	// 3. 双向流调用
	bidirectionalChat(client)
}

// 1. 客户端流
func clientStreamUpload(client pb.StreamServiceClient) {
	stream, _ := client.ClientStreamUpload(context.Background())
	for i := 0; i < 5; i++ {
		_ = stream.Send(&pb.UploadRequest{Data: "Message " + strconv.Itoa(i+1)})
		time.Sleep(500 * time.Millisecond)
	}
	resp, _ := stream.CloseAndRecv()
	log.Println("Upload Response:", resp.Status)
}

// 2. 服务端流
func serverStreamDownload(client pb.StreamServiceClient) {
	req := &pb.DownloadRequest{Count: 5}
	stream, _ := client.ServerStreamDownload(context.Background(), req)

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		log.Println("Received:", resp.Data)
	}
}

// 3. 双向流
func bidirectionalChat(client pb.StreamServiceClient) {
	stream, _ := client.BidirectionalChat(context.Background())

	// 发送消息
	go func() {
		for i := 0; i < 3; i++ {
			item := i + 1
			_ = stream.Send(&pb.ChatMessage{User: "Client", Message: "Hello " + strconv.Itoa(item)})
			time.Sleep(1 * time.Second)
		}
		_ = stream.CloseSend()
	}()

	// 接收消息
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		log.Println("Received:", resp.Message)
	}
}
