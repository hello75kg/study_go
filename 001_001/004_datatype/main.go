package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int8
	var b int16
	var c int32
	var d int64

	var ua uint8
	var ub uint8
	var uc uint8
	var ud uint8

	//int是动态类型，根据操作系统决定，
	//32位操作系统是int32
	//64位操作系统是int64
	var i int

	b = int16(a)

	var f1 float32 //3.4e38
	var f2 float64 //1.8e308
	fmt.Println(a, b, c, d, ua, ub, uc, ud, i, b, f1, f2)

	// byte 底层是 uint8 的别名
	// byte 用来存放字符(ascii)，就是char
	var b1 byte
	b1 = 'a'
	fmt.Println(b1)         //97
	fmt.Printf("%c \n", b1) //a
	b1 = b1 + 1
	fmt.Println(b1)         // 98
	fmt.Printf("%c \n", b1) //b

	//rune 底层是 int32 的别名 type rune int32
	//也是字符，byte存不下中文等符号
	var r1 rune
	r1 = '王'
	fmt.Println(r1)
	fmt.Printf("%c \n", r1)

	//字符串
	var name string
	name = "王王王"
	fmt.Println(name)

	//类型转换
	//定义一个类型 type TI int32
	type TI int32
	var a32 int32
	var ti32 TI
	a32 = 9
	ti32 = TI(a32)
	fmt.Println(ti32)

	//字符串转换
	var istr string = "123"
	var intstr int
	intstr, err := strconv.Atoi(istr)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(intstr)

	istr = strconv.Itoa(intstr)
	fmt.Println(istr)

	//字符串转float32，bool
	f, err := strconv.ParseFloat("3.14", 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f)

	//parseInt 可以转其他进制
	ii, err := strconv.ParseInt("11", 16, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ii)

	//1 为true 0 为false，其他false，"true"为true，"false"转为false
	pb, err := strconv.ParseBool("1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pb)

	//相反，其他类型转字符串，用format
	fb := strconv.FormatBool(true)
	fmt.Println(fb)

	fstr := strconv.FormatFloat(3.14, 'f', -1, 64)
	fmt.Println(fstr)
	fstr = strconv.FormatFloat(3.14123442345453, 'E', -1, 64)
	fmt.Println(fstr)

	//进制转换
	fii := strconv.FormatInt(10, 16)
	fmt.Println(fii)

}
