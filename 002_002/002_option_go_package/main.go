package main

func main() {
	// go_package
	// 用于指定 Go 代码生成时的包路径和包名，确保在不同的 Go 模块或项目中正确导入 Protobuf 生成的代码
	// 由 两部分 组成：
	// 1.	导入路径（可选）：common/proto/v1
	// 		生成的 .pb.go 文件会放到该路径下。
	// 		代码中 import 这个包时会使用该路径。
	// 2.	包名（必选）：hello
	// 		生成的 .pb.go 文件中的 package 关键字会使用该名称。
	// 		在 Go 代码中使用时，应该 import "gcommon/proto/v1"，然后用 hello.xxx{} 访问消息类型。

}
