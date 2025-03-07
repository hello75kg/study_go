package main

import "fmt"

func main() {
	//运算法和表达式

	//字符串长度
	//字符长度为1，中文字符长度是3
	name := "abc王王王"
	fmt.Println(len(name))

	bytes := []rune(name)
	fmt.Println(len(bytes))

	//转义符
	cName := "abcd\"abc"
	fmt.Println(cName)
	cName = `abc'abc"abc`
	fmt.Println(cName)

}
