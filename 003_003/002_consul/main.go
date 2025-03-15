package main

func main() {
	// 管理页面默认端口 127.0.0.1:8500
	// 1.服务注册与发现:	允许服务动态注册，并支持 DNS / HTTP API 查询
	// 2.健康检查:		定期检查服务是否可用，不可用时自动剔除
	// 3.负载均衡:		通过 Consul DNS 解析实现简单的负载均衡
	// 4.Key-Value 存储:	存储分布式配置，如数据库连接、API Key 等
	// 5.多数据中心支持:	可在多个数据中心同步数据，实现跨地域容灾
	//
	// go get github.com/hashicorp/consul/api
	//
	// 使用 Consul DNS 进行服务发现
	// Consul 允许你使用 DNS 解析来获取服务实例
	// nslookup user-service.service.consul
	// dig @127.0.0.1 -p 8600 consul.service.consul SRV
	//
	// kv存储：
	// 	consul kv put app/config/db "mysql://root:password@localhost:3306"
	//	consul kv get app/config/db
	//
	// 多数据中心：
	// consul join -wan 192.168.1.2
	//

	// 服务注册：
	// PUT 请求： 127.0.0.1:8500/v1/agent/service/register
	// {
	//    "Name":"wshop-web",
	//    "ID":"wshop-web",
	//    "Tags":["wshop","wang","chen","web"],
	//    "Port":50051
	// }
	// 服务注销：
	// PUT 请求：127.0.0.1:8500/v1/agent/service/deregister/wshop-web
	//
}
