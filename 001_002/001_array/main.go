package main

import "fmt"

func main() {
	//数组、切片、map、list
	//数组：定义：var arr [count]int

	var courses3 [3]string
	var courses4 [4]string

	fmt.Printf("%T \n", courses3)
	fmt.Printf("%T \n", courses4)

	//打印类型，一个是[3]string,一个是[4]string，是两种类型，不能用一个给另一个赋值
	//courses4 = courses3（错误）

	//[]string 和 [3]string 不同
	//一个是切片，一个是固定元素个数的数组
	var cs1 []string
	var cs2 [3]string
	cs2[0] = "go"
	cs2[1] = "gin"
	cs2[2] = "grpc"
	fmt.Printf("cs2: %T \n", cs2)
	fmt.Printf("cs1: %T \n", cs1)
	fmt.Printf("%v \n", cs1)
	for k, v := range cs2 {
		fmt.Println(k, v)
	}

	//初始化
	var css1 [3]string = [3]string{"aaa", "bbb", "ccc"}
	var css2 = [3]string{"aaa", "bbb", "ccc"}
	css3 := [3]string{"aaa", "bbb", "ccc"}
	css4 := [3]string{1: "bbb"}
	css5 := [...]string{"aaa", "bbb", "ccc"} // ...编译执行是数组长度自动为3
	fmt.Println(css1, css2, css3, css4, css5)
	for k, v := range css4 {
		fmt.Println(k, v)
	}

	//比较
	//[2]string 和 [3]string
	//比较编译时会报错，类型不匹配，而不是不相等
	sarr1 := [2]string{"aaa", "bbb"}
	sarr2 := [3]string{"aaa", "bbb", "ccc"}
	fmt.Println(sarr2)
	//sarr1 == sarr2 （invalid operation: sarr1 == sarr2 (mismatched types [2]string and [3]string)）
	sarr3 := [2]string{"aaa", "bbb"}
	fmt.Println(sarr1 == sarr3) //内容按顺序都相等的数组相等

	//多维数组
	var csInfo [3][4]string
	csInfo[0] = [4]string{"aaa", "bbb", "ccc", "ddd"}
	csInfo[1] = [4]string{"aaa", "bbb", "ccc"}
	csInfo[2] = [4]string{"aaa", "bbb", "ccc"}
	csInfo[2][3] = "ddd"
	fmt.Println(csInfo)
	for _, v := range csInfo {
		for _, v2 := range v {
			fmt.Print(v2, "\t")
		}
		fmt.Println()
	}

}
