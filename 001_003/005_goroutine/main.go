package main

import (
	"fmt"
	"time"
)

// 多进程和多线程切换资源消耗较大
// 协程是用户态，切换消耗小
// go协程内存占用小（2k），切换快，go没有多线程，只有协程可用 goroutine
// GMP模型

func asynPrint() {
	for true {
		fmt.Println("111111")
		time.Sleep(time.Second * 30)
	}
}

// main主协程，主协程退了协程也会退出
func main() {
	//go asynPrint()
	//fmt.Println("main")

	//for 循环变量 i 在整个循环过程中是 同一个变量，多个 goroutine 共享这个变量。
	//当 goroutine 真正执行时，i 可能已经改变，导致输出结果不符合预期
	//可能会发现多个 goroutine 输出相同的 i 值（通常接近 1000000），而不是从 0 到 999999 逐个打印
	for i := 0; i < 1000000; i++ {
		go func() {
			fmt.Println(i)
		}()
	}

	//这样，每个 goroutine 都会 捕获不同的 i 值，输出才符合预期
	for i := 100; i < 200; i++ {
		num := i
		go func() {
			fmt.Println(num)
		}()
	}

	//或者，把i传进去
	for i := 200; i < 300; i++ {
		go func(num int) {
			fmt.Println(num)
		}(i)
	}

	time.Sleep(5 * time.Second)
}
