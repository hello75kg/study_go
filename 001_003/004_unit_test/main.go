package main

func main() {
	//单元测试
	//go test
	//按照一定约定和组织的测试代码驱动程序
	//包目录中以_test.go为后缀的源码文件都会被go test运行到,并且不会被build打包到可执行文件中
	//4种类型
	//Test开头的 功能测试 Benchmark开头的 性能测试 example 模糊测试

	//跳过耗时的单元测试,用 go test -short

	//性能测试Benchmark开头的 go test -bench= ".*"

}
