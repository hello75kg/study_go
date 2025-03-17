# Go 的 Redsync 分布式锁原理

## 1. Redsync 简介
Redsync 是基于 Redis 实现的分布式锁库，适用于 Go 语言。它利用 Redis 的 SET 命令及其 `NX` 和 `PX` 选项，实现高效的分布式锁。

## 2. 基本原理
Redsync 采用 **Redlock 算法**，其核心思想是：
- 在多个独立的 Redis 实例上 **同时加锁**，以提高容错能力。
- 使用 **随机唯一值** 作为锁的标识，确保锁的安全性。
- 通过 **超时时间** 自动释放锁，避免死锁。

## 3. Redlock 算法流程
Redlock 主要包括以下步骤：

1. **获取当前时间戳**（毫秒级）。
2. **依次尝试在多个 Redis 节点上加锁**：
    - 使用 `SET resource_name unique_value NX PX lock_ttl`。
    - `NX`：如果键不存在，则创建（确保互斥性）。
    - `PX lock_ttl`：设置锁的过期时间（防止死锁）。
3. **计算加锁时间**：
    - 计算成功获取锁的 Redis 实例数量，若过半（一般为 3/5），则认为加锁成功。
4. **检查加锁时间**：
    - 若获取锁的时间小于 `TTL/2`，则锁成功，否则释放锁（防止网络延迟导致锁超时）。
5. **执行业务逻辑**。
6. **释放锁**：
    - 只有加锁时存入的唯一标识与锁值匹配时，才能执行 `DEL` 释放锁，防止误删。

## 4. Redsync 代码示例
```go
package main
import (
    "fmt"
    "github.com/go-redsync/redsync/v4"
    "github.com/go-redsync/redsync/v4/redis/goredis/v8"
    "github.com/redis/go-redis/v9"
    "context"
    "time"
)

func main() {
    // 创建 Redis 客户端
    client := redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
    })

    // 适配 Redsync
    pool := goredis.NewPool(client)
    rs := redsync.New(pool)

    // 创建互斥锁
    mutex := rs.NewMutex("my-lock")

    // 尝试加锁
    ctx := context.Background()
    if err := mutex.LockContext(ctx); err != nil {
        fmt.Println("获取锁失败:", err)
        return
    }

    fmt.Println("获取锁成功")

    // 模拟业务处理
    time.Sleep(2 * time.Second)

    // 释放锁
    if ok, err := mutex.UnlockContext(ctx); !ok || err != nil {
        fmt.Println("释放锁失败:", err)
    } else {
        fmt.Println("释放锁成功")
    }
}
```

## 5. Redlock 的优缺点

优点
-	高可用性：即使部分 **Redis** 节点失效，仍可加锁。
-	防止死锁：通过 **TTL** 机制自动释放锁。
-	避免误删：使用 **唯一标识** 校验释放锁。

缺点
-	**网络延迟影响**：加锁时间过长可能导致锁失败。
-	**一致性问题**：理论上，某些极端情况下可能发生锁误抢。

## 6. 适用场景
   -	分布式任务调度（防止任务重复执行）。
   -	限流控制（防止超量并发）。
   -	资源竞争（多个进程争抢共享资源）。

Redsync 提供了一种可靠的分布式锁方案，但在高并发和高可靠性需求下，仍需要结合业务需求进行优化。