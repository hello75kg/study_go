package main

import "fmt"

func myPrint(items ...interface{}) {
	for k, v := range items {
		fmt.Println(k, v)
	}
}

func main() {
	var data = []interface{}{1, "2", 3.4}
	myPrint(data...)

	//var data2 =[]string{"1", "2", "3"}
	//报错，因为[]string和[]interface底层结构不一样，不能隐式转换
	//myPrint(data2...)

	//可以显式地把 []string 转换成 []interface{}
	var data2 = []string{"1", "2", "3"}
	var data2Interface = make([]interface{}, len(data2))
	for i, v := range data2 {
		data2Interface[i] = v
	}
	myPrint(data2Interface...) // 这样就不会报错
}
