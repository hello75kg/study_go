package main

import "net/http"
import (
	"github.com/gin-gonic/gin"
)

// go.mod
// go的modules文件是自动维护的
// import 后sync会自动从github上下载相关的包
// go modules 添加依赖，自动包管理
// go 命令： go get ; go mod help ; go install
// go mod tidy //整理包，下载引用新包、删除不用的包引用
// go get -u 升级包版本
// go get -u patch 升级到最新的修订版本
// go项目名和modules名字不一样时，可以用在go.mod中用replace替换
// 如：gin: 【import   "github.com/gin-gonic/gin"】
// go.mod下的go.sum是用来校验包里面的文件的
// 命令行执行 go env 查看go环境配置
// 设置国内镜像源：go env -w GOPROXY='https://goproxy.cn,direct'

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	// http://127.0.0.1:8080/ping
}
