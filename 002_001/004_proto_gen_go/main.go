package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"

	"studyProject/002_001/004_proto_gen_go/proto"
)

type Hello struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Courses []string `json:"courses"`
}

func main() {
	// grpc
	// grpc 是一个高性能、开源、通用的rpc框架，面向移动和HTTP/2设计
	// protocol buffer 是 google 的一种轻量、高效的数据存储结构，目前pb3

	// 工具
	// 下载 protoc，加入环境变量
	// protoc-gen-go
	// // 执行 go get github.com/golang/protobuf/protoc-gen-go
	// 执行 go get google.golang.org/protobuf/protoc-gen-go
	// brew install protoc-gen-go
	//
	// 用 *.proto 文件生成 go 代码
	// protoc --go_out=. --go-grpc_out=.  helloworld.proto

	req := __.HelloRequest{
		Name: "wang",
	}
	marshal, _ := proto.Marshal(&req)
	fmt.Println(marshal)
	fmt.Println(string(marshal))

	jsonStruct := Hello{
		Name: "wang",
	}
	jsonres, _ := json.Marshal(&jsonStruct)
	fmt.Println(string(jsonres))

	proto2 := __.HelloRequest{
		Name:    "wang",
		Age:     18,
		Courses: []string{"111", "222", "333"},
	}
	marshal, _ = proto.Marshal(&proto2)
	fmt.Println(string(marshal))

	jsonStruct2 := Hello{
		Name:    "wang",
		Age:     18,
		Courses: []string{"111", "222", "333"},
	}
	jsonres, _ = json.Marshal(&jsonStruct2)
	fmt.Println(string(jsonres))

	newReq := __.HelloRequest{}
	_ = proto.Unmarshal(marshal, &newReq)
	fmt.Println(newReq.Name, newReq.Age, newReq.Courses)

}
