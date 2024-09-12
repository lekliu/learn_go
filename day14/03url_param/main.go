package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// API 参数

func main() {
	// 创建一个默认的路由引擎
	// 默认使用了2个中间件Logger(), Recovery,
	r := gin.Default()
	r.GET("/welcome", func(c *gin.Context) {
		name := c.DefaultQuery("name", "Jack")
		c.String(http.StatusOK, fmt.Sprintf("Hello, %s", name))
	})
	r.Run(":8000")
}

//http://127.0.0.1:8000/welcome?name=Naccy
//out:  Hello, Naccy
