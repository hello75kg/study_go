package main

import (
	"fmt"
	"strings"
)

func add(a, b int) int {
	return a + b
}

func add2(a, b interface{}) int {
	return a.(int) + b.(int)
}
func add3(a, b interface{}) interface{} {
	switch a.(type) {
	case int:
		return a.(int) + b.(int)
	case float64:
		return a.(float64) + b.(float64)
	case string:
		return a.(string) + b.(string)
	default:
		panic("unknown type")
	}
}

func main() {
	a := 1
	b := 2
	fmt.Println(add(a, b))
	fmt.Println(add2(a, b))
	fmt.Println(add3(1.1, 2.2))
	strA := "hello"
	strB := " world"
	res := add3(strA, strB)
	fmt.Println(res)
	fmt.Printf("%T\n", res)
	//返回的res是interface类型，需要用断言转成string
	split := strings.Split(res.(string), " ")
	fmt.Println(split)

}
