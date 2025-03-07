package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//格式化输出
	username := "王王王"
	address := "北京"
	age := 19

	fmt.Printf("username :%s, address:%s, age:%d \n", username, address, age)

	spf := fmt.Sprintf("username :%s, address:%s, age:%d ", username, address, age)
	fmt.Println(spf)

	//%v 省略格式打印 数组[1,2,3]
	//%#v 保留go定义代码输出
	//%T 输出类型
	var arr = []int{1, 2, 3}
	fmt.Printf("%v \n", arr)  //[1 2 3]
	fmt.Printf("%#v \n", arr) //[]int{1, 2, 3}
	fmt.Printf("%T \n", arr)  //[]int

	//字符串拼接
	//高性能字符串拼接 strings.Builder
	var builder strings.Builder
	builder.WriteString("用户名：")
	builder.WriteString(username)
	builder.WriteString(" 年龄：")
	builder.WriteString(strconv.Itoa(age))
	builder.WriteString(" 地址：")
	builder.WriteString(address)

	bs := builder.String()
	fmt.Println(bs)

}
