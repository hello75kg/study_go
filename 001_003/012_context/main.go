package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // 监听 context 取消信号
			fmt.Println("Worker stopped")
			return
		default:
			fmt.Println("Worker running...")
			time.Sleep(time.Second)
		}
	}
}

func process(ctx context.Context) {
	traceID := ctx.Value("trace_id") // 获取 trace_id
	fmt.Println("Processing request with trace_id:", traceID)
}

func main() {
	// context
	// context 是go用于控制 goroutine 生命周期的标准库
	// 主要用于
	// 超时控制：设置请求的最大执行时间
	// 取消操作：多个goroutine之间共享取消信号
	// 传递请求范围的元数据：如trace_id、用户认证信息

	// 常见方式
	// context.Background()	最顶层的context，通常用于main函数、初始化、测试
	// context.TODO()	未来可能要替换context的占位符
	// context.WithCancel(parent)	创建可取消的context
	// context.WithTimeout(parent,time.Duration)	设定超时时间，超时后自动取消
	// context.WithValue(parent,key,value)	传递元数据（如 trace_id）

	// context.WithCancel(parent)	创建可取消的context
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx)              // 启动 goroutine
	time.Sleep(3 * time.Second) // 运行 3 秒
	cancel()                    // 发送取消信号
	time.Sleep(1 * time.Second) // 等待 worker 退出

	// context.WithTimeout(parent,time.Duration)	设定超时时间，超时后自动取消
	ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second) // 3 秒超时
	defer cancel()                                                         // 确保资源释放
	go worker(ctx)
	time.Sleep(5 * time.Second) // 主进程等待

	// context.WithValue 传递元数据（如 trace_id）
	ctx = context.WithValue(context.Background(), "trace_id", "123456") // 传递 trace_id
	process(ctx)

	// context.TODO()	未来可能要替换context的占位符
	// TODO() 不具备超时控制和取消功能,只是一个占位符，不会触发ctx.Done()
	// worker() 会 一直运行，因为 ctx.Done() 永远不会触发
	ctx = context.TODO() // 这里暂时用 TODO，之后可以换成 WithCancel 或 WithTimeout
	go worker(ctx)
	time.Sleep(3 * time.Second)
	fmt.Println("Main function done.")
}
