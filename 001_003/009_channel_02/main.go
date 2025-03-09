package main

import (
	"fmt"
	"time"
)

func Producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i
	}
	close(out)
}

func Consumer(in <-chan int) {
	for num := range in {
		fmt.Println(num)
	}
}

func main() {
	// 单向 channel
	var chnl chan int = make(chan int, 3)
	// 只能发送数据的channel
	var chnlSend chan<- int = chnl
	// 只能接受数据的channel
	var chnlRecv <-chan int = chnl
	chnlSend <- 1
	fmt.Println(<-chnlRecv)

	fmt.Println("-----------")
	go Producer(chnlSend)
	go Consumer(chnlRecv)

	time.Sleep(time.Second)

}
