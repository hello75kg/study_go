package main

func main() {
	// grpc 状态码
	// 在 gRPC 中，状态码（gRPC Status Codes）用于表示 RPC 调用的结果，类似于 HTTP 状态码。
	// 它们定义在 google.golang.org/grpc/codes 包中。

	// 在 gRPC 服务器端，使用 status.Errorf(codes.Code, "message") 返回状态码
	// 在 gRPC 客户端，可以使用 status.Code(err) 来解析错误

	// 状态码	名称					说明							HTTP 等效
	// 0		OK					请求成功						200 OK
	// 1		CANCELLED			请求被客户端取消				499 Client Closed Request
	// 2		UNKNOWN				未知错误						500 Internal Server Error
	// 3		INVALID_ARGUMENT	参数无效						400 Bad Request
	// 4		DEADLINE_EXCEEDED	请求超时						504 Gateway Timeout
	// 5		NOT_FOUND			资源未找到					404 Not Found
	// 6		ALREADY_EXISTS		资源已存在					409 Conflict
	// 7		PERMISSION_DENIED	权限不足						403 Forbidden
	// 8		RESOURCE_EXHAUSTED	资源耗尽（限流、配额不足）		429 Too Many Requests
	// 9		FAILED_PRECONDITION	违反前置条件					412 Precondition Failed
	// 10		ABORTED				操作中止（通常用于并发冲突）		409 Conflict
	// 11		OUT_OF_RANGE		数值超出范围					400 Bad Request
	// 12		UNIMPLEMENTED		服务器未实现方法				501 Not Implemented
	// 13		INTERNAL			服务器内部错误				500 Internal Server Error
	// 14		UNAVAILABLE			服务器不可用（宕机/网络故障）	503 Service Unavailable
	// 15		DATA_LOSS			数据丢失或损坏				500 Internal Server Error
	// 16		UNAUTHENTICATED		认证失败						401 Unauthorized
	//
}
