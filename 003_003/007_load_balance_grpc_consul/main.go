package main

func main() {
	// 负载均衡
	// 1. 集中式 load balance
	//		1.1.web-api 层通过 consul dns 查询所有service地址列表
	//		1.2.通过一个第三方负债均衡器（lvs、 nginx），所有请求都走负载均衡器获得实际请求的某一个service地址
	// 2. 进程内 load balance（grpc使用这种方式）
	//		2.1.service 向注册中心注册服务
	//		2.2.web-api 层起一个协程，向注册中心获取一批连接，维护一个连接池
	//		2.3.请求从连接池中获取一个连接，去请求 service 服务
	//		2.4.每种语言需要一套自己的 sdk
	// 3. 独立进程 load balance，中和 1 和 2
	//		3.1.在每台 web-api 上部署一个独立进程的负载均衡器
	//		3.2.需要多维护一套单独的逻辑，还需要监控进程可用（watchdog?）

	// 负载均衡算法
	// 1. 轮询：按顺序轮流分配请求到后端服务器上，无视服务器状态和负载
	// 2. 随机
	// 3.源地址哈希：根据客户端请求 ip 地址哈希值转成数字，按服务器数量取模
	// 4.加权轮询
	// 5.加权随机
	// 6.最小接连数：连接数量少的服务器优先

	// 库
	// go get github.com/mbobakov/grpc-consul-resolver

}
