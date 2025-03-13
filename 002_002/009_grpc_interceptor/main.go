package main

func main() {
	// grpc 拦截器
	// 用于RPC方法执行前后插入自定义逻辑
	// 日志记录、身份认证、限流、错误处理、指标收集
	// 1. Unary 拦截器（一元拦截器）：单次请求-响应的rpc应用
	// 2. Stream 拦截器 （流式拦截器）：流式rpc调用（Client Streaming / Server Streaming / Bidirectional）

	// 1. Unary 拦截器
	// 		服务端：UnaryInterceptor
	// 		客户端：WithUnaryInterceptor

	// 2. Stream 拦截器
	// 		StreamInterceptor

	// 服务端拦截器链，从外到内执行：
	// func chainUnaryInterceptors(interceptors ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	// 	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// 			// 递归调用所有拦截器
	// 			var chainHandler grpc.UnaryHandler = handler
	// 			for i := len(interceptors) - 1; i >= 0; i-- {
	// 				currInterceptor := interceptors[i]
	// 				chainHandler = func(currentCtx context.Context, currentReq interface{}) (interface{}, error) {
	// 				return currInterceptor(currentCtx, currentReq, info, chainHandler)
	// 			}
	// 		}
	// 		return chainHandler(ctx, req)
	// 	}
	// }
	//
	// grpcServer := grpc.NewServer(
	// 	grpc.UnaryInterceptor(chainUnaryInterceptors(loggingInterceptor, authInterceptor)),
	// )
}
