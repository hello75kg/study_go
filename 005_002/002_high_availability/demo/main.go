package main

import (
	"fmt"
	"log"
	"time"

	"github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/alibaba/sentinel-golang/core/circuitbreaker"
)

// 初始化基于错误数的熔断规则
func initErrorCountBreaker() {

	_, err := circuitbreaker.LoadRules([]*circuitbreaker.Rule{
		{
			Resource:         "error_count_api",
			Strategy:         circuitbreaker.ErrorCount,
			Threshold:        10,    // 10 次错误触发熔断
			MinRequestAmount: 5,     // 需要至少 5 个请求
			StatIntervalMs:   10000, // 10s 窗口统计
			RetryTimeoutMs:   5000,  // 熔断 5s 后尝试恢复
		},
	})
	if err != nil {
		log.Fatalf("熔断规则加载失败: %v", err)
	}
}

func callAPI() {
	e, err := api.Entry("error_count_api", api.WithTrafficType(base.Inbound))
	if err != nil {
		fmt.Println("请求被熔断:", err)
		return
	}
	defer e.Exit()

	// 模拟失败请求
	fmt.Println("请求失败")
	e.Exit(base.WithError(fmt.Errorf("service error")))
}

func main() {
	api.InitDefault()
	initErrorCountBreaker()

	for i := 0; i < 20; i++ {
		go callAPI()
		time.Sleep(500 * time.Millisecond)
	}

	time.Sleep(20 * time.Second)
}
