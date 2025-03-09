package main

import (
	"fmt"
	"sync"
)

func main() {
	// 用channel实现交叉打印数字和字母
	// 1A2B3C4D5E6F6G7H...

	var chNum = make(chan struct{})
	var chChar = make(chan struct{})
	var chDone = make(chan struct{})
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		num := 1
		for {
			select {
			case <-chNum:
				fmt.Print(num)
				num++
				chChar <- struct{}{}
			case <-chDone:
				close(chNum)
				return
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		ch := 'A'
		for {
			select {
			case <-chChar:
				fmt.Printf("%c", ch)
				ch++
				if ch > 'Z' {
					close(chChar)
					close(chDone)
					return
				}
				chNum <- struct{}{}
			}
		}
	}()
	chNum <- struct{}{}
	wg.Wait()
}
