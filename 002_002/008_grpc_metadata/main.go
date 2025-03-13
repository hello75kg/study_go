package main

func main() {
	// gRPC 的 Metadata（元数据）
	// 主要用于 在客户端和服务器之间传递额外的上下文信息，类似于 HTTP 头部（HTTP Headers）。
	// 可以用于：
	//		身份验证（如传递 Token）
	//		请求追踪（Trace ID、Request ID）
	//		自定义信息（如自定义 Header）
	//
	// Metadata 的键值格式
	//		键名：区分大小写，一般以 小写字母 + "-" 分隔 方式命名（如 authorization）。
	//		值：可以是字符串、二进制数据（以 -bin 结尾）。
	//		普通键：普通字符串，如 authorization: Wang token123。
	//		二进制键：键名 必须以 -bin 结尾，值是 Base64 编码的二进制数据。
	//
	//
}
