package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	pb "studyProject/002_001/006_grpc_stream/proto"
)

type server struct {
	pb.UnimplementedStreamServiceServer
}

// 1. 客户端流：接收多个消息后返回响应
func (s *server) ClientStreamUpload(stream pb.StreamService_ClientStreamUploadServer) error {
	var count int
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.UploadResponse{Status: fmt.Sprintf("Received %d messages", count)})
		}
		if err != nil {
			return err
		}
		count++
		log.Println("Received:", req.Data)
	}
}

// 2. 服务端流：向客户端持续发送数据
func (s *server) ServerStreamDownload(req *pb.DownloadRequest, stream pb.StreamService_ServerStreamDownloadServer) error {
	for i := 0; i < int(req.Count); i++ {
		_ = stream.Send(&pb.DownloadResponse{Data: fmt.Sprintf("Message %d", i+1)})
		time.Sleep(time.Second) // 模拟延迟
	}
	return nil
}

// 3. 双向流：实时聊天
func (s *server) BidirectionalChat(stream pb.StreamService_BidirectionalChatServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Received from %s: %s", req.User, req.Message)

		// 回复客户端
		_ = stream.Send(&pb.ChatMessage{User: "Server", Message: "Echo: " + req.Message})
	}
}

func main() {
	listener, _ := net.Listen("tcp", ":50051")
	grpcServer := grpc.NewServer()
	pb.RegisterStreamServiceServer(grpcServer, &server{})
	fmt.Println("gRPC server running on port 50051")
	_ = grpcServer.Serve(listener)
}
