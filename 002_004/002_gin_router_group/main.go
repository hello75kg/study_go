package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/goods/list", func(c *gin.Context) {})
	router.GET("/goods/1", func(c *gin.Context) {})
	router.GET("/goods/add", func(c *gin.Context) {})
	// 分组：
	goodsGroup := router.Group("/goods")
	goodsGroup.GET("/list", func(c *gin.Context) {})
	goodsGroup.POST("/1", func(c *gin.Context) {})
	goodsGroup.PUT("/add", func(c *gin.Context) {})

	// 简单的路由组: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", loginEndpoint)
		v1.POST("/submit", submitEndpoint)
		v1.POST("/read", readEndpoint)
	}

	// 简单的路由组: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", loginEndpoint)
		v2.POST("/submit", submitEndpoint)
		v2.POST("/read", readEndpoint)
	}

	router.Run(":8080")
}

func readEndpoint(context *gin.Context) {

}

func submitEndpoint(context *gin.Context) {

}

func loginEndpoint(context *gin.Context) {

}
