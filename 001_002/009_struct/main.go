package main

import "fmt"

type Person struct {
	name    string
	age     int
	address string
}

type Animal struct {
	name string
	age  int
}
type Cat struct {
	Animal
	food string
}

func (animal Animal) wow() string {
	return animal.name
}
func (cat Cat) wow() string {
	return "cat:" + cat.name
}

func (cat *Cat) setAge(age int) {
	cat.age = age
}

func main() {
	//没有指定结构体中的字段，必须顺序赋值所有字段
	var p0 = Person{"chen", 19, "GZ"}
	fmt.Println(p0)
	//指定了字段名赋值，可以只赋值需要的字段
	var p1 Person = Person{
		name: "wang",
		age:  20,
	}
	fmt.Println(p1)

	persons := []Person{p0, p1, {name: "六"}}
	fmt.Println(persons)

	p0.age = 38
	fmt.Println(p0.age)

	//匿名结构体
	addr := struct {
		province string
		city     string
	}{
		"FJ",
		"FZ",
	}
	fmt.Println(addr.province, addr.city)

	//结构体嵌套（类似继承）
	var cat Cat = Cat{
		Animal{name: "mmm", age: 1},
		"mice",
	}
	cat.name = "mm"
	cat.age = 20
	cat.food = "fish"
	fmt.Println(cat)

	//给结构体绑定方法
	//不用改变时传值，要改变时传引用
	//func (animal Animal) wow() string{
	//	return animal.name
	//}
	//func (animal Animal) wow() string{
	//	return animal.name
	//}
	fmt.Println(cat.wow())
	cat.setAge(10)
	fmt.Println(cat.age)
}
