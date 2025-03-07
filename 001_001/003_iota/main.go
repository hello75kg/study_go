package main

import "fmt"

func a() (int, bool) {
	return 0, false
}

func main() {
	//匿名变量
	var _ int
	_, ok := a()
	ok = true
	if ok {
		fmt.Println(1)
	}
	//作用域
	var name = "name"
	if ok {
		name = "ok_name"
		fmt.Println(name)
	}
	fmt.Println(name)

	//iota 特殊常量 ，可以被编译器修改的常量
	const (
		//自动赋值
		ERR1 = iota + 1
		ERR2
		ERR3 = "haha"
		ERR4
		ERR5
		ERR6 = iota
		ERR7
	)
	const (
		E0 = iota
		E1
	)
	fmt.Println(ERR1, ERR2, ERR3, ERR4, ERR5, ERR6, ERR7)

}
