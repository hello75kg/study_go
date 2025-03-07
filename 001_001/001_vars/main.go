package main

import "fmt"

var name = "wang"
var age = 1
var ok bool

var (
	name2 = "wang"
	age2  = 2
	ok2   bool
)

func main() {
	//1.先定义后使用
	//2.必须有类型
	//3.类型定义后不能改变
	//4.局部变量定义了不使用是不行的，全局变量可以
	//var name int
	//var age = 1
	age := 1
	fmt.Println(age)

	var u1, u2, u3 string
	var uu1, uu2, uu3 = "w1", 1, "w3"
	fmt.Println(u1, u2, u3)
	fmt.Println(uu1, uu2, uu3)
	//go是静态类型，类型和赋值类型一致
	//简洁变量定义不能用于全局 i:=1
	//变量有0值 默认值

}
