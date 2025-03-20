package middleware

import (
	"fmt"
	"net/http"

	"github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/gin-gonic/gin"
)

// Sentinel 限流中间件
func SentinelMiddleware(resource string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 进入 Sentinel 资源
		e, err := api.Entry(resource, api.WithTrafficType(base.Inbound))
		if err != nil {
			// 请求被限流
			fmt.Println("Sentinel 限流了")
			c.JSON(http.StatusTooManyRequests, gin.H{"message": "请求过多，请稍后再试"})
			c.Abort()
			return
		}
		defer e.Exit()

		// 继续处理请求
		fmt.Println("正常处理业务逻辑...")
		c.Next()
	}
}
