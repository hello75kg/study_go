package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	go func() {
		r.Run(":8080")
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	// 收到退出信号时的后续处理
	// 如果是 kill -9 强制关闭，就没办法继续执行了
	fmt.Println("server关闭中...")
	fmt.Println("服务注销...")
}
