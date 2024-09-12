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
	// 路由组1，处理GET请求
	v1 := r.Group("/v1")
	// {} 是书写规范
	{
		v1.GET("/login", login)
		v1.GET("/submit", submit)
	}

	v2 := r.Group("/v2")
	{
		v2.POST("/login", login)
		v2.POST("/submit", submit)
	}

	r.Run(":8000")
}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "Jack")
	c.String(http.StatusOK, fmt.Sprintf("Hello, %s", name))
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "Naccy")
	c.String(http.StatusOK, fmt.Sprintf("Hello, %s", name))
}

//C:\>curl http://127.0.0.1:8000/v2/submit -X POST
//Hello, Naccy
