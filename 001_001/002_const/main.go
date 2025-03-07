package main

func main() {
	//常量，定义时赋值，不能修改
	const PI float32 = 3.14
	const PI2 = 3.14
	const (
		PI3 = 3.14
		PI4 = 3.14
		PI5 = 3.14
	)
	// 定义一组的时候会，如果没有申明类型和值，会延用上一个
	const (
		a = 14
		b
		c = "s1"
		d
	)
}
