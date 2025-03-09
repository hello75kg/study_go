package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	// mutex
	// 锁不能复制
	var mtx sync.Mutex
	var waitGroup sync.WaitGroup
	num := 0
	for i := 1; i <= 1000; i++ {
		waitGroup.Add(1) // 必须放在go外面，协程启动前，否则可能还有协程没开始运行导致没有add主协程就退出了
		go func() {
			defer waitGroup.Done()
			mtx.Lock()
			defer mtx.Unlock()
			num = num + 1
		}()
	}
	waitGroup.Wait()
	fmt.Println(num)

	// 原子操作(简单操作)
	// atomic
	var num2 int32 = 0
	for i := 1; i <= 1000; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			atomic.AddInt32(&num2, 1)
		}()
	}
	waitGroup.Wait()
	fmt.Println(num2)

	// RWMutex
	// 读写锁，一般情况下读操作多余锁操作，读与读之间可以不用锁，只有读和写之间才需要锁
	// 读和读不冲突，只有读和写才有冲突
	// 读遇到读锁，可读
	// 读遇到写锁，不可读
	// 写遇到读锁，不可写
	var num3 int32 = 0
	var rwMutex sync.RWMutex

	for i := 1; i <= 5; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			for {
				rwMutex.Lock()
				fmt.Println("get Write lock")
				time.Sleep(time.Second * 5000)
				num3 = num3 + 1
				rwMutex.Unlock()
			}
		}()
	}

	for i := 0; i < 5; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			for {
				rwMutex.RLock()
				fmt.Println("get Read lock")
				time.Sleep(time.Millisecond * 500)
				rwMutex.RUnlock()
			}
		}()
	}
	waitGroup.Wait()

}
