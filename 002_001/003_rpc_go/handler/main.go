package handler

const HelloServiceName = "handler/HelloService"

type HelloService struct{}

func (sh *HelloService) Hello(req string, reply *string) error {
	// 返回值是通过修改传入的reply
	*reply = "Hello " + req
	return nil
}
