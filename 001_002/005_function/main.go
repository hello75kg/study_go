package main

import (
	"fmt"
)

// 函数参数是值传递
func add(a int, b int) (int, error) {
	sum := a + b
	return sum, nil
}

// 返回类型里也可以加变量名
func add2(a int, b int) (sum int, err error) {
	sum = a + b
	return sum, nil
}

// 在返回类型里定义了变量名时，可以直接return，不用再显式指定返回
// 如果return返回的不是定义的变量名sum，会把实际返回为值赋值给定义的变量名，再返回
func add3(a int, b int) (sum int, err error) {
	sum = a + b
	return
}

// 参数是切片，用...int
func add4(items ...int) (sum int, err error) {
	for _, item := range items {
		sum += item
	}
	return
}

func cal(a int, b int, f func(c int, d int) int) int {
	res := f(a, b)
	return res
}

func main() {
	//普通函数
	sum, _ := add(1, 2)
	fmt.Println(sum)
	sum, _ = add2(1, 2)
	fmt.Println(sum)
	sum, _ = add3(1, 2)
	fmt.Println(sum)
	sum, _ = add4(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(sum)

	//函数可以作为变量传递
	varFunc := add4
	sum, _ = varFunc(1, 2, 3, 4, 5)
	fmt.Println(sum)

	//匿名函数
	sum = cal(1, 1, func(c int, d int) int {
		return c + d
	})
	fmt.Println(sum)
	sum = cal(1, 1, func(c int, d int) int {
		return c - d
	})
	fmt.Println(sum)

	fmt.Println()
	//闭包
	//在一个函数a中再定一个函数b,每次调用函数b来操作函数a中定义的局部变量
	//例如：有个函数每个调用返回的值都比上一次大1
	af := autoIncr()
	fmt.Println(af())
	fmt.Println(af())
	fmt.Println(af())
	fmt.Println(af())

}

func autoIncr() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
