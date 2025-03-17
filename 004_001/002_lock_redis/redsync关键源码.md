# Go 的 Redsync 关键源码解读

## 1. Redsync 简介
Redsync 是一个基于 Redis 的分布式锁实现，主要用于 Go 语言项目。它使用多个 Redis 实例来提高锁的可靠性，符合 **Redlock 算法**。

## 2. 核心组件
Redsync 的核心组件如下：
- **Mutex**：表示一个分布式锁实例
- **Redsync**：管理多个 Redis 实例
- **Pool**：Redis 连接池，支持多个 Redis 实例

## 3. 关键源码解析

### 3.1 创建 Redsync 实例
```go
func New(pools ...Pool) *Redsync {
    return &Redsync{
        pools: pools,
    }
}
```
`New` 方法接收多个 `Pool` 作为参数，用于初始化 Redsync。

### 3.2 生成 Mutex（锁）
```go
func (r *Redsync) NewMutex(name string, options ...Option) *Mutex { 
    m := &Mutex{
        name: name,
        pools: r.pools,
        ...
    }
    return m
}
```
该方法创建一个 `Mutex` 对象，并关联多个 Redis 实例。

### 3.3 获取锁（Lock）
```go
func (m *Mutex) Lock() (bool, error) {
    value := randomValue()
    expiry := time.Now().Add(m.expiry)

    n := 0
    for _, pool := range m.pools {
        conn := pool.Get()
        defer conn.Close()
        
        res, err := conn.Do("SET", m.name, value, "NX", "PX", m.expiry.Milliseconds())
        if err == nil && res == "OK" {
            n++
        }
    }

    if n >= quorum(len(m.pools)) {
        return true, nil
    }
    
    m.Unlock()
    return false, ErrLockFailed
}
```
- 生成随机值 `value`，用于唯一标识当前锁
- 遍历多个 Redis 实例，使用 `SET NX PX` 方式尝试加锁
- 如果多数（`quorum`）Redis 实例成功加锁，则返回成功
- 若加锁失败，则释放已加的锁（`Unlock`）

### 3.4 释放锁（Unlock）
```go
func (m *Mutex) Unlock() bool {
    n := 0
    for _, pool := range m.pools {
    conn := pool.Get()
    defer conn.Close()

    script := `
        if redis.call("GET", KEYS[1]) == ARGV[1] then
            return redis.call("DEL", KEYS[1])
        else
            return 0
        end`
        
        res, err := redis.Int(conn.Do("EVAL", script, 1, m.name, m.value))
        if err == nil && res == 1 {
            n++
        }
    }

    return n >= quorum(len(m.pools))
}
```
- 通过 `EVAL` 执行 Lua 脚本，确保只有持有锁的客户端能释放
- 计算释放成功的 Redis 实例数量，若达到 `quorum` 则解锁成功

## 4. 结论
Redsync 通过 **多实例写入 + Quorum 机制** 确保高可用，采用 **Lua 脚本** 确保锁的原子性，是 Go 语言实现分布式锁的优秀方案。
