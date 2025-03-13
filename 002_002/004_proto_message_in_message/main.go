package main

import (
	"fmt"
	"studyProject/002_002/004_proto_message_in_message/proto"
)

func main() {
	res := proto.HelloReply{}
	res.Message = "Hello World"
	// 嵌套类型需要加前缀
	res.Data = make([]*proto.HelloReply_Result, 1)
	fmt.Println(res)
}
