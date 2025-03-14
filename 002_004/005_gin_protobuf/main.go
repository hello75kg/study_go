package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"studyProject/002_004/005_gin_protobuf/proto"
)

func main() {
	router := gin.Default()
	router.GET("/moreJson", moreJson)
	router.GET("/someProtoBuf", someProtoBuf)

	router.Run(":8080")
}

func someProtoBuf(c *gin.Context) {
	c.ProtoBuf(http.StatusOK, &proto.Hello{
		Name: "wang",
	})
}

func moreJson(c *gin.Context) {
	var msg struct {
		Name    string `json:"user"` // 返回中将Name重命名为user
		Message string
		Number  int
	}
	msg.Name = "wang"
	msg.Message = "message"
	msg.Number = 20
	// c.JSON(http.StatusOK, msg)
	c.PureJSON(http.StatusOK, msg)
}
