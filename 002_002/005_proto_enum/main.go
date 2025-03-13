package main

import (
	"fmt"
	"studyProject/002_002/005_proto_enum/proto"
)

func main() {
	hello := proto.Hello{
		Gender: proto.Gender_MALE,
	}
	fmt.Println(hello)
}
