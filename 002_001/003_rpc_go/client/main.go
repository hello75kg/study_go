package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"studyProject/002_001/003_rpc_go/client_proxy"

	"studyProject/002_001/003_rpc_go/handler"
)

func main() {
	conn, _ := rpc.Dial("tcp", "localhost:1234")
	var reply string
	// _ = conn.Call("HelloService.Hello", "world", &reply)
	_ = conn.Call(handler.HelloServiceName+".Hello", "world", &reply)
	fmt.Println(reply)

	// json-rpc
	conn2, _ := net.Dial("tcp", "localhost:2234")
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn2))
	var reply2 string
	// _ = client.Call("HelloService.Hello", "world", &reply2)
	_ = client.Call(handler.HelloServiceName+".Hello", "world", &reply2)
	fmt.Println(reply2)

	// 解偶
	serviceClient := client_proxy.NewHelloServiceClient("tcp", "localhost:1234")
	var reply3 string
	_ = serviceClient.Hello("world", &reply3)
	fmt.Println(reply3)
}
