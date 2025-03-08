package main

import (
	"fmt"
	"strconv"
)

type myInt3 int

func (mi myInt3) toString() string {
	return strconv.Itoa(int(mi))
}

func main() {
	//type
	//1. 定义结构体
	//2. 定义接口
	//3. 定义类型别名
	//4. 定义类型
	//5. 类型判断

	//定义类型别名，myInt会在编译时改为int
	type myInt = int
	var a myInt = 1
	sum := a + 1
	fmt.Printf("%T \n", a)
	fmt.Println(sum)

	//定义类型，myInt2的实际类型时main.myInt2，而不是改为int，所以和int类型不同，不能直接相加，可以强转成int后操作
	type myInt2 int
	var b myInt2 = 1
	fmt.Printf("%T \n", b)
	sum = int(b) + 1
	fmt.Println(sum)

	//定义类型可以给已知的类型增加方法，别名不可以
	var c myInt3 = 1
	fmt.Println(c.toString())

	//类型判断
	var d interface{} = "d interface string"
	fmt.Printf("%T \n", d)
	switch d.(type) {
	case string:
		fmt.Println("string")

	}

}
