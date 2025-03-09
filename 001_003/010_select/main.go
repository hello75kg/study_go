package main

import (
	"fmt"
	"time"
)

func main() {
	// select
	// 主要作用于多个channel
	// select 语句在 Go 语言中用于处理多个 channel 的读写操作，
	// 类似于 switch，但它的作用是 监听多个 channel 的数据流动，并执行 最先满足条件的 case 语句

	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	// select 监听 ch1、ch2 的 接收操作 和 ch3 的 发送操作。
	// 哪个 channel 先准备好，select 就执行对应的 case。
	// 如果 多个 channel 都可用，select 会随机选择一个执行（防止某个 channel 饿死）。
	// 如果所有 case 都阻塞，则 select 也会阻塞，除非有 default 语句。
	select {
	case msg := <-ch1:
		fmt.Println("Received from ch1:", msg)
	case msg := <-ch2:
		fmt.Println("Received from ch2:", msg)
	case ch3 <- "hello":
		fmt.Println("Sent to ch3")
	default:
		fmt.Println("No channel is ready")
	}

	// 监听多个 channel（防止阻塞）
	// 模拟不同的通道有不同的消息延迟
	go func() {
		time.Sleep(time.Second * 1)
		ch1 <- "Message from ch1"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		ch2 <- "Message from ch2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case msg := <-ch2:
			fmt.Println(msg)
		}
	}

	// 处理 channel 发送（防止 channel 关闭后写入 panic）
	// 通过 done channel 让 goroutine 安全退出，避免 ch 关闭后仍然写入导致 panic
	ch := make(chan string, 1)
	done := make(chan bool)

	go func() {
		for {
			select {
			case ch <- "hello":
				fmt.Println("Message sent")
			case <-done:
				fmt.Println("Stopping goroutine")
				return
			}
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(3 * time.Second)
	close(done) // 关闭 done，通知 goroutine 停止
	time.Sleep(1 * time.Second)

	// 超时处理
	// select 还能用于 超时处理，防止某个 channel 迟迟不返回，导致程序卡住
	go func() {
		time.Sleep(3 * time.Second)
		ch <- "result"
	}()

	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)
	// 如果 ch 在 2 秒内没有数据，就执行 case <-time.After(2 * time.Second)，避免 goroutine 卡死
	case <-time.After(2 * time.Second): // 超时 2 秒
		fmt.Println("Timeout, no response")
	}
}
