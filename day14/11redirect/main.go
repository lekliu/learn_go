package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HTML渲染

func main() {
	// 创建一个默认的路由引擎
	// 默认使用了2个中间件Logger(), Recovery,
	r := gin.Default()

	r.GET("/redirect", func(c *gin.Context) {
		//支持内部和外重定向
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
	})

	r.Run(":8000")
}

//http://127.0.0.1:8000/index
