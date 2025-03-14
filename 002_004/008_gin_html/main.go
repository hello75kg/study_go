package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 官网： https://gin-gonic.com/
// 安装： go get -u github.com/gin-gonic/gin

func main() {
	r := gin.Default()
	// r.LoadHTMLFiles("templates/index.html")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Hello Gin",
		})
	})
	r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
