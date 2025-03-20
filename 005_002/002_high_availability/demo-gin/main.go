package main

import (
	"fmt"
	"log"
	"net/http"
	"studyProject/005_002/002_high_availability/demo-gin/middleware"

	"github.com/gin-gonic/gin"

	"github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/flow"
)

// 初始化 Sentinel
func initSentinel() {
	// 初始化 Sentinel
	err := api.InitDefault()
	if err != nil {
		log.Fatalf("Sentinel 初始化失败: %v", err)
	}

	// 定义限流规则
	_, err = flow.LoadRules([]*flow.Rule{
		{
			Resource:               "GET:/api/limit", // 资源名称（建议使用 URL 作为资源）
			Threshold:              5,                // QPS 阈值，超过 5 请求会被限流
			TokenCalculateStrategy: flow.Direct,      // 直接限流策略
			ControlBehavior:        flow.Reject,      // 拒绝请求
		},
	})
	if err != nil {
		log.Fatalf("限流规则加载失败: %v", err)
	}

	fmt.Println("Sentinel 限流规则加载成功")
}

func main() {
	// 初始化 Sentinel
	initSentinel()

	// 创建 Gin 服务器
	r := gin.Default()

	// 限流路由
	r.GET("/api/limit", middleware.SentinelMiddleware("GET:/api/limit"), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "请求成功"})
	})

	// 启动服务器
	r.Run(":18080")
}
