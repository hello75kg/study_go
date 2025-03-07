package main

import (
	"fmt"
	"strings"
)

func main() {

	a := "hello"
	b := "world"
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a > b)

	a = "abcd网e-fg王"
	b = "bcd"
	//包含
	cs := strings.Contains(a, b)
	fmt.Println(cs)

	//统计
	cc := strings.Count(a, b)
	fmt.Println(cc)

	//分割
	st := strings.Split(a, "-")
	fmt.Println(st)

	//包含前缀
	hp := strings.HasPrefix(a, "abc")
	fmt.Println(hp)

	//包含后缀
	hsf := strings.HasSuffix(a, "fg")
	fmt.Println(hsf)

	//查询子串位置，一个中文占3个长度
	ir := strings.IndexRune(a, '王')
	fmt.Println(ir)

	//替换
	srpls := strings.Replace(a, "a", "b", 1)
	fmt.Println(srpls)

	//大小写转换
	upr := strings.ToUpper(a)
	fmt.Println(upr)

	//Trim去掉首尾特殊字符,左边TrimLeft,右边TrimRight
	trim := strings.Trim(a, "ab")
	fmt.Println(trim)

}
