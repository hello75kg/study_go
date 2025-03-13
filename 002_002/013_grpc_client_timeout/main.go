package main

// 客户端超时，如果服务器超过 2 秒未响应，客户端将收到 DeadlineExceeded 错误。
// func callWithTimeout(client pb.MyServiceClient) {
// 		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
// 		defer cancel()
//
// 		resp, err := client.MyMethod(ctx, &pb.MyRequest{})
// 		if err != nil {
// 		st, _ := status.FromError(err)
// 		log.Printf("gRPC call failed: %v, Message: %s", st.Code(), st.Message())
// 			return
// 		}
//
// 		log.Println("Response:", resp)
// }

// 客户端流超时
// func streamWithTimeout(client pb.MyServiceClient) {
//		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
//		defer cancel()
//
//		stream, err := client.StreamMethod(ctx, &pb.MyRequest{})
//		if err != nil {
//			log.Println("Stream failed:", err)
//			return
//		}
//
//		for {
//			resp, err := stream.Recv()
//			if err == io.EOF {
//				break
//			}
//			if err != nil {
//				log.Println("Stream error:", err)
//				break
//			}
//			log.Println("Received:", resp)
//		}
// }

// 服务端，服务端可以通过 ctx.Done() 检测超时，并提前返回错误。
// func (s *server) MyMethod(ctx context.Context, req *pb.MyRequest) (*pb.MyResponse, error) {
//		// 模拟一个耗时任务（5 秒）
//		select {
//		case <-time.After(5 * time.Second):
//			return &pb.MyResponse{Message: "Success"}, nil
//		case <-ctx.Done(): // 如果客户端超时或取消（如果客户端超时，ctx.Done() 会被触发，服务器可以提前结束任务，避免资源浪费。）
//			return nil, status.Errorf(codes.DeadlineExceeded, "Request timed out")
//		}
// }

func main() {
	// grpc
	// 在 gRPC 中，超时机制用于限制 RPC 调用的最大执行时间，防止服务端长时间无响应，影响系统稳定性。
	// gRPC 超时由客户端设置，服务端接收并遵守。

	//	1.	客户端 context.WithTimeout 或 context.WithDeadline
	//		适用于 单次 RPC 请求 的超时控制
	//	2.	服务端 context.Context 读取超时时间
	//		适用于 服务器端业务逻辑 需要感知超时

}
