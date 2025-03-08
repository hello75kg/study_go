package main

import (
	"errors"
	"fmt"
)

func funcErr() (int, error) {
	return 0, errors.New("an error")
}

func funcPanic() (int, error) {
	panic("a panic")
	return 0, errors.New("a panic error")
}

func main() {
	//go错误处理的理念
	//go语言函数中以返回值return一个错误的方式处理错误
	a, err := funcErr()
	fmt.Println(a)
	fmt.Println(err)

	defer func() {
		re := recover()
		if re != nil {
			fmt.Println("a recover1")
		}
	}()

	//panic
	//panic函数会导致程序退出
	//一般是服务启动过程中，必要的资源获取失败时调用
	//如数据库连接失败，读取不到必要的配置文件等
	b, err := funcPanic()
	fmt.Println(b)
	fmt.Println(err)

	//recover
	//上面panic后，程序立即中断，main退出，上面panic前的defer生效，触发recover1执行，而下面的代码因为main已经退出了就不再执行
	//所以这里输出recover1
	defer func() {
		re := recover()
		if re != nil {
			fmt.Println("a recover2")
		}
	}()
}
