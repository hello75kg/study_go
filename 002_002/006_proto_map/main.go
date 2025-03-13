package main

import (
	"fmt"
	"studyProject/002_002/006_proto_map/proto"
)

func main() {
	hello := proto.Hello{
		Mp: map[string]string{
			"hello": "world",
			"a":     "b",
		},
	}
	fmt.Println(hello)
}
