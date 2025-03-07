package main

import (
	"fmt"
)

func main() {
	//if
	age := 18
	day := 232
	if age >= 18 && day >= 211 {
		fmt.Println(age)
	} else if age == 18 {
		fmt.Println(day)
	} else {
		fmt.Println("...")
	}

	//for
	//go没有while
	ok := true
	i := 0
	for ok {
		i++
		if i > 10 {
			break
		}
		fmt.Println(i)
	}

	for {
		//time.Sleep(1 * time.Second)
		fmt.Println(i)
		i++
		if i > 16 {
			break
		}
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 打印九九乘法表
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d x %d = %d \t", j, i, i*j)
		}
		fmt.Println()
	}

	// for range
	// 字符串、数组、切片、map、channel
	name := "hello world 王"
	for k, v := range name {
		fmt.Printf("%d,%c \n", k, v)
	}

	nameRune := []rune(name)
	for k, v := range nameRune {
		fmt.Println(k, v)
	}

	//continue
	//break
	//goto

	// switch
	switch age {
	case 18:
		fmt.Println(age)
	case 19, 22, 32:
		fmt.Println(age)
	case 20:
		fmt.Println(age)
	default:
		fmt.Println(age)
	}

}
