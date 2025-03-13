package main

func main() {
	// grpc 验证器
	// 安装 proto-gen-validate
	// https://github.com/bufbuild/protoc-gen-validate

	// go install github.com/envoyproxy/protoc-gen-validate@latest
	// protoc --go_out=. --go-grpc_out=. --validate_out=lang=go:. stream.proto

	//
	// message ValidateRequest {
	//	 	string email = 1 [(validate.rules).string.email = true]; // 必须是有效的 Email
	// 		int32 age = 2 [(validate.rules).int32 = {gte: 18, lte: 100}]; // 18-100 之间
	// }

	// server 端实现
	// func (s *server) ValidateExample(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	//		// 使用 PGV 进行参数验证
	//		if err := validate.Validate(req); err != nil {
	//			return nil, fmt.Errorf("validation failed: %v", err)
	//		}
	//		return &pb.ValidateResponse{Message: "Valid request!"}, nil
	// }

}
