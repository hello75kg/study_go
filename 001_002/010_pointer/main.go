package main

import "fmt"

type Person struct {
	name string
}

func setName(person *Person) {
	person.name = "wow"

}

func main() {
	//函数传的事值，直接传修改没用
	//&取地址
	//*取地址指向的数据
	var person Person
	setName(&person)
	fmt.Printf("%p \n", &person)
	fmt.Println(person.name)

	//go语言限制了指针运算，比如指针地址+1往前移一位，go不支持
	//如果要运算要使用unsafe包

	//指针初始化
	//var p *int
	//fmt.Println(p)
	var pi int = 10
	var pip = &pi
	fmt.Println(*pip)

	//变量申明没有赋值会有初始值，比如0
	//指针变量声明了没有赋值则为nil，不能直接使用
	//推荐用new
	ps := &Person{}
	var ps2 Person
	ps3 := &ps2
	var ps4 = new(Person)
	fmt.Println(ps)
	fmt.Println(*ps3)
	fmt.Println(*ps4)

	//交换两个int
	a := 1
	b := 2
	func(a, b *int) {
		//传进来时的a和b的地址
		//要交换两个值，需要用*获取地址指向的值做交换
		*a, *b = *b, *a
	}(&a, &b)
	fmt.Println(a, b)
	//nil
	//不同数据的0值不一样
	//bool false
	//numbers 0
	//string “”
	// nil只代表以下6种类型空值，不含结构体，// Type must be a pointer, channel, func, interface, map, or slice type
	//pointer nil
	//slice nil
	//map nil
	//channel nil
	//interface nil
	//function nil

}
