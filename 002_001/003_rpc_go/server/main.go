package main

import (
	"io"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"studyProject/002_001/003_rpc_go/server_proxy"

	"studyProject/002_001/003_rpc_go/handler"
)

// 放到handler里
// type HelloService struct{}
// func (sh *HelloService) Hello(req string, reply *string) error {
// 	// 返回值是通过修改传入的reply
// 	*reply = "Hello " + req
// 	return nil
// }

func main() {
	// 实例化server
	listener, _ := net.Listen("tcp", ":1234")
	listener2, _ := net.Listen("tcp", ":2234")
	// 注册rpc服务
	// _ = rpc.RegisterName("HelloService", &handler.HelloService{})
	// _ = rpc.RegisterName(handler.HelloServiceName, &handler.HelloService{})
	_ = server_proxy.RegisterHelloService(&handler.HelloService{})
	// 情动服务端口接收请求
	for {
		accept, _ := listener.Accept()
		go rpc.ServeConn(accept)
		// json-rpc
		accept2, _ := listener2.Accept()
		go rpc.ServeCodec(jsonrpc.NewServerCodec(accept2))

		// http-jsonrpc
		// _ = rpc.RegisterName("HelloService", &HelloService{})
		// _ = rpc.RegisterName(handler.HelloServiceName, &handler.HelloService{})
		_ = server_proxy.RegisterHelloService(&handler.HelloService{})
		http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
			var conn io.ReadWriteCloser = struct {
				io.Writer
				io.ReadCloser
			}{
				ReadCloser: r.Body,
				Writer:     w,
			}
			_ = rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
		})
		_ = http.ListenAndServe(":3234", nil)
	}

	// 序列化/反序列化? server_proxy 和 client_proxy 能否自动生成？:
	// protobuf + gRPC
}
