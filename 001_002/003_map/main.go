package main

import "fmt"

func main() {
	// map 是一个key-value无序集合，查询快

	var cMap map[string]string
	fmt.Println(cMap)
	//map想要直接通过key赋值必须要先初始化
	//如上面定义的cMap， cMap["aaa"]="AAA"会报错
	//cMap["aaa"] = "AAA"//(assignment to entry in nil map)

	var cMap2 = map[string]string{}
	//这样就不会报错了，{}实现了初始化
	cMap2["aaa"] = "AAA"
	fmt.Println(cMap2)

	cMap3 := map[string]string{
		"go":   "gogoog",
		"grpc": "grpcgrpcgrpc",
		"gin":  "gingingin",
	}
	fmt.Println(cMap3)

	//遍历
	fmt.Println("----------")
	for k, v := range cMap3 {
		fmt.Println(k, v)
	}
	// map是无序的，遍历时每次打印顺序可能会不一样
	for k := range cMap3 {
		fmt.Println(k, cMap3[k])
	}

	fmt.Println("----------")

	//获取数据、删除元素
	d := cMap3["go"]
	fmt.Println(d)

	delete(cMap3, "go")
	fmt.Println(cMap3)
	// map不是线程安全的

}
