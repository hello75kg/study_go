package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 官网： https://gin-gonic.com/
// 安装： go get -u github.com/gin-gonic/gin

func main() {
	r := gin.Default()
	r.Use(MyLogger(), TokenRequired())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}

func TokenRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		for k, v := range c.Request.Header {
			if k == "X-Token" {
				token = v[0]
				break
			}
		}
		if token == "wang" {
			c.JSON(http.StatusOK, gin.H{})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			// 中断请求后续执行的逻辑
			c.Abort()
		}
		c.Next()
	}
}

func MyLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Set("example", "123456")
		// 让原本要执行的代码继续执行完
		c.Next()

		since := time.Since(start)
		fmt.Println("time:", since)
		status := c.Writer.Status()
		fmt.Println("status:", status)
	}
}
