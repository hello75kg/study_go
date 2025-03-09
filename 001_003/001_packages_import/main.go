package main

import (
	"fmt"
	_ "studyProject/001_003/001_packages_import/001_demo"    //用下划线做别名，不使用包，但是需要执行init()方法做初始化
	demo "studyProject/001_003/001_packages_import/001_demo" //别名
	//. "studyProject/001_003/001_packages_import/001_demo" //用.可以把包里面的内容当作是当前包里的，就不用xxx.来调用，不推荐
)

func main() {
	//package
	//用来组织源码，是多个go源码的组合，是代码复用的基础
	//每个源码的开头必须定义package
	//同一个目录下的源码必须同一个package
	//同一个目录下的源码可以互相引用，不用import，直接用【包名加点】访问
	//不同目录下的源码引用，import的地址从根路径的包开始
	//不同的包之间，只用调用首字母大写的结构
	var p = demo.People{
		Name: "aaa",
	}
	fmt.Println(p)

	var m = Main2{
		Name: "aaa",
	}
	fmt.Println(m)
}
