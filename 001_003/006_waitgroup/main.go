package main

import (
	"fmt"
	"sync"
)

func main() {
	//WaitGroup
	//用于主协程中等待子协程处理完成

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			fmt.Println(num)
		}(i)
	}

	wg.Wait()
}
