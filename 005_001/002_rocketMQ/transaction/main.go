package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"time"
)

// TransactionListener 事务监听器
type TransactionListener struct{}

// ExecuteLocalTransaction 执行本地事务逻辑
func (t *TransactionListener) ExecuteLocalTransaction(msg *primitive.Message) primitive.LocalTransactionState {
	fmt.Println("执行本地事务...")
	// // 模拟事务状态未知，等待回查的情况
	// return primitive.UnknowState
	// 模拟本地事务执行逻辑
	success := doLocalTransaction()
	if success {
		fmt.Println("本地事务执行成功，提交消息")
		return primitive.CommitMessageState
	} else {
		fmt.Println("本地事务执行失败，回滚消息")
		return primitive.RollbackMessageState
	}
}

// CheckLocalTransaction RocketMQ 事务回查
func (t *TransactionListener) CheckLocalTransaction(msg *primitive.MessageExt) primitive.LocalTransactionState {
	fmt.Println("RocketMQ 回查本地事务...")

	// 这里应该查询本地事务状态（如数据库状态）
	success := checkLocalTransactionStatus()
	if success {
		fmt.Println("本地事务已成功，提交消息")
		return primitive.CommitMessageState
	} else {
		fmt.Println("本地事务失败，回滚消息")
		return primitive.RollbackMessageState
	}
}

// 模拟本地事务执行
func doLocalTransaction() bool {
	time.Sleep(2 * time.Second) // 模拟事务处理时间
	return true                 // 事务成功返回 true，失败返回 false
}

// 模拟事务状态检查
func checkLocalTransactionStatus() bool {
	return true
}

func main() {
	// 创建 RocketMQ 事务生产者
	p, _ := rocketmq.NewTransactionProducer(
		&TransactionListener{},
		producer.WithNameServer([]string{"192.168.0.249:9876"}), // RocketMQ 地址
		// producer.WithVIPChannel(false),
	)

	err := p.Start()
	if err != nil {
		fmt.Println("启动 Producer 失败:", err)
		return
	}

	// 发送事务消息
	msg := &primitive.Message{
		Topic: "hellormq",
		Body:  []byte("Hello RocketMQ Transaction!"),
	}

	// 发送事务消息
	res, err := p.SendMessageInTransaction(context.Background(), msg)
	if err != nil {
		fmt.Println("发送事务消息失败:", err)
	} else {
		fmt.Printf("发送事务消息成功: %s\n", res.String())
	}

	// 运行一段时间观察事务回查
	time.Sleep(10 * time.Minute)

	// 关闭生产者
	err = p.Shutdown()
	if err != nil {
		fmt.Println("关闭 Producer 失败:", err)
	}
}
