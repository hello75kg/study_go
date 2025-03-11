package server_proxy

import (
	"net/rpc"
	"studyProject/002_001/003_rpc_go/handler"
)

type HelloServicer interface {
	Hello(req string, reply *string) error
}

func RegisterHelloService(srv HelloServicer) error {
	return rpc.RegisterName(handler.HelloServiceName, srv)

}
