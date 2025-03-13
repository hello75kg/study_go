package main

import (
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"studyProject/002_002/007_proto_timestamp/proto"
	"time"
)

func main() {
	hello := proto.Hello{
		Time: timestamppb.New(time.Now()),
	}
	fmt.Println(hello)
}
