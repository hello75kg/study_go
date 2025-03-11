package client_proxy

import (
	"net/rpc"
	"studyProject/002_001/003_rpc_go/handler"
)

type HelloServiceSub struct {
	*rpc.Client
}

func NewHelloServiceClient(protcal string, address string) HelloServiceSub {
	conn, _ := rpc.Dial(protcal, address)
	return HelloServiceSub{conn}
}

func (c *HelloServiceSub) Hello(args string, reply *string) error {
	_ = c.Client.Call(handler.HelloServiceName+".Hello", args, reply)
	return nil
}
