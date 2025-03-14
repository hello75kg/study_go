package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person struct {
	ID   int    `uri:"id" binding:"required"`
	Name string `uri:"name" binding:"required"`
}

func main() {
	router := gin.Default()

	// url里有变量，用冒号：
	goodsGroup := router.Group("/goods")
	goodsGroup.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{"id": id})
	})
	goodsGroup.GET("/:id/:action", func(c *gin.Context) {
		id := c.Param("id")
		action := c.Param("action")
		c.JSON(http.StatusOK, gin.H{"id": id, "action": action})
	})
	router.GET("/:name/:id", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBindUri(&person); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"name": person.Name, "id": person.ID})
	})

	router.Run(":8080")
}
