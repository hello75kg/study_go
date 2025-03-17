# Redis 分布式锁详细原理

## 1. 什么是分布式锁？
分布式锁是一种用于在分布式系统中确保多个进程或线程对共享资源的互斥访问的机制。常见应用场景包括：
- 订单去重
- 限制并发操作
- 任务调度

在单机环境下，可以使用 `synchronized`（Java）或 `mutex`（Go）等机制实现锁。但在分布式环境下，由于多个实例之间不能共享内存，需要借助 Redis 等存储系统实现分布式锁。

---

## 2. Redis 分布式锁的实现方式

Redis 提供了多种方式来实现分布式锁，其中最常见的是使用 `SET` 命令：

```bash
SET lock_key unique_value NX EX 10
```

### 2.1 SET NX EX 方式
- `lock_key`：锁的 key，表示需要加锁的资源
- `unique_value`：唯一值，防止误删锁（通常使用 UUID）
- `NX`（Not Exists）：如果 key 不存在，则设置 key，保证互斥性
- `EX 10`（Expire 10s）：设置 10 秒超时，防止死锁

#### 示例代码（Go 语言）
```go
package main

import (
    "context"
    "fmt"
    "github.com/redis/go-redis/v9"
    "time"
)

var rdb = redis.NewClient(&redis.Options{
    Addr: "localhost:6379",
})

func acquireLock(ctx context.Context, key string, value string, expiration time.Duration) bool {
    result, err := rdb.SetNX(ctx, key, value, expiration).Result()
    if err != nil {
        fmt.Println("Error acquiring lock:", err)
        return false
    }
    return result
}

func main() {
    ctx := context.Background()
    locked := acquireLock(ctx, "my_lock", "unique_id_123", 10*time.Second)
    if locked {
        fmt.Println("Lock acquired!")
    } else {
        fmt.Println("Failed to acquire lock")
    }
}
```

---

## 3. 释放锁的安全性
由于锁的 `EX` 过期时间可能会提前释放，多个客户端可能会误删他人的锁。因此，删除锁时必须保证：
- 只有加锁的客户端才能释放锁
- 使用 Lua 脚本保证操作的原子性

```bash
luaScript := `
if redis.call("GET", KEYS[1]) == ARGV[1] then
    return redis.call("DEL", KEYS[1])
else
    return 0
end
`
```

### 示例代码（Go 语言）
```go
func releaseLock(ctx context.Context, key string, value string) bool {
    luaScript := `
    if redis.call("GET", KEYS[1]) == ARGV[1] then
        return redis.call("DEL", KEYS[1])
    else
        return 0
    end
    `
    result, err := rdb.Eval(ctx, luaScript, []string{key}, value).Result()
    if err != nil {
        fmt.Println("Error releasing lock:", err)
        return false
    }
    return result.(int64) == 1
}
```

---

## 4. 锁的续期（Watchdog 机制）
锁的 `EX` 过期时间如果太短，可能导致锁意外释放；太长又可能造成死锁。因此，**需要续期机制**：
1. 后台守护线程（Watchdog）定期续期
2. 只有持有锁的进程才可以续期

示例代码：
```go
func renewLock(ctx context.Context, key string, value string, expiration time.Duration) {
    ticker := time.NewTicker(expiration / 2)
    for range ticker.C {
        luaScript := `
        if redis.call("GET", KEYS[1]) == ARGV[1] then
            return redis.call("PEXPIRE", KEYS[1], ARGV[2])
        else
            return 0
        end
        `
        result, _ := rdb.Eval(ctx, luaScript, []string{key}, value, expiration.Milliseconds()).Result()
        if result.(int64) == 0 {
            fmt.Println("Lock expired, stopping renewal")
            return
        }
    }
}
```

---

## 5. RedLock 算法（提高可靠性）
单个 Redis 节点的锁存在以下问题：
- **主从复制延迟**：如果主节点崩溃，锁可能会在从节点不同步的情况下丢失
- **数据丢失**：如果 Redis 以 `AOF` 方式持久化但未及时写入磁盘，锁信息可能丢失

**RedLock** 方案：
- 部署 **5 个 Redis 实例**
- 获取 **大多数（N/2+1）** 实例上的锁
- 计算 **加锁时间**，确保小于锁的有效期
- 释放时 **删除所有实例上的锁**

示例流程：
1. 依次向 `5` 个 Redis 实例请求 `SET lock_key unique_value NX PX 30000`
2. 统计成功获取锁的 Redis 实例个数（至少 `3` 个）
3. 计算锁获取总耗时，若耗时过长（接近 `30s`），则放弃锁
4. 业务执行完成后，删除 `5` 个 Redis 实例上的锁

示例代码：
```go
// RedLock 客户端
```

---

## 6. 总结
| 方案            | 特点 |
|---------------|------------------------------------------------|
| `SET NX EX`   | 最常见，适用于单实例 Redis，不适用于高可用场景 |
| Lua 释放锁    | 确保删除时锁仍然属于自己 |
| Watchdog 续期 | 防止锁超时提前释放 |
| RedLock       | 适用于分布式环境，提高可靠性，但性能开销更大 |

Redis 分布式锁适用于 **短时锁**，对于长时间锁，可以考虑 **ZooKeeper** 方案。