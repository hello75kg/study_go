package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// channel
	// 不要通过共享内存来通信
	// 而是通过通信来共享内存
	//
	// 语法糖
	// <-
	//
	// 读取已经关闭且无数据的channel，会返回该类型的0值
	//
	// 应用场景
	// 1. 消息传递、消息过滤
	// 2. 信号广播
	// 3. 事件订阅和广播
	// 4. 任务分发
	// 5. 结果汇总
	// 6. 并发控制
	// 7. 同步和异步
	// ...

	// 缓存空间
	var msg chan string
	msg = make(chan string, 2) // 初始化缓存空间大小如果不够用，会阻塞
	msg <- "hello"
	msg <- "world"
	fmt.Println(<-msg)
	fmt.Println(<-msg)

	// 有缓存空间
	// 适用于生产者和消费者之间通信
	// 无缓存空间
	// 适用于通知，一个协程要第一时间知道另一个协程是否处理完成

	// 无缓存区
	var wg sync.WaitGroup
	var chnl chan string
	chnl = make(chan string, 0)

	wg.Add(1)
	go func(chnl chan string) {
		defer wg.Done()
		defer func() {
			// channel已经关闭
			if r := recover(); r != nil {
				fmt.Println("Recovered in f", r)
			}
		}()
		for {
			chnl <- "hello 1 "
			chnl <- "world 1 "
			// time.Sleep(time.Second * 1)
		}
	}(chnl)

	wg.Add(1)
	go func(chnl chan string) {
		defer wg.Done()
		// for {
		//	data := <-chnl
		//	fmt.Println(data)
		// }
		// 也可以这样写
		for data := range chnl {
			fmt.Println(data)
			time.Sleep(time.Second)
		}
		fmt.Println("done")
	}(chnl)

	time.Sleep(time.Second * 2)
	// 关闭的channel不能再继续传值，但是还能取值
	close(chnl)
	print(<-chnl) // 输出字符串的 0 值，字符串的 0 值为""
	print(<-chnl)
	print(<-chnl)
	wg.Wait()

}
