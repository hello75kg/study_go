package main

import (
	"fmt"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	goredislib "github.com/redis/go-redis/v9"
	"sync"
)

func main() {
	//
	// 基于 redis 的分布式锁
	//
	// Create a pool with go-redis (or redigo) which is the pool redisync will
	// use while communicating with Redis. This can also be any pool that
	// implements the `redis.Pool` interface.
	client := goredislib.NewClient(&goredislib.Options{
		Addr: "192.168.0.249:6379",
	})
	pool := goredis.NewPool(client) // or, pool := redigo.NewPool(...)

	// Create an instance of redisync to be used to obtain a mutual exclusion
	// lock.
	rs := redsync.New(pool)

	// Obtain a new mutex by using the same name for all instances wanting the
	// same lock.
	gNum := 10
	mutexname := "my-global-mutex"
	var wg sync.WaitGroup
	for i := 0; i < gNum; i++ {
		iNum := i
		wg.Add(1)
		go func(taskId int) {
			defer wg.Done()
			mutex := rs.NewMutex(mutexname)
			fmt.Printf("任务 %d，开始获取锁 \n", taskId)
			if err := mutex.Lock(); err != nil {
				fmt.Printf("任务 %d，获取锁失败 \n", taskId)
				return
			}
			fmt.Printf("任务 %d，获取锁成功 \n", taskId)
			fmt.Printf("任务 %d，开始释放锁 \n", taskId)
			if _, err := mutex.Unlock(); err != nil {
				fmt.Printf("任务 %d，释放锁失败 \n", taskId)
				return
			}
			fmt.Printf("任务 %d，释放锁成功 \n", taskId)

		}(iNum)
	}
	wg.Wait()

	// Obtain a lock for our given mutex. After this is successful, no one else
	// can obtain the same lock (the same mutex name) until we unlock it.
	// if err := mutex.Lock(); err != nil {
	// 	panic(err)
	// }

	// Do your work that requires the lock.

	// Release the lock so other processes or threads can obtain a lock.
	// if ok, err := mutex.Unlock(); !ok || err != nil {
	// 	panic("unlock failed")
	// }
}
