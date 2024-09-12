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
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		c.String(http.StatusOK, fmt.Sprintf(" %s is %s \n", name, action))
	})
	r.Run(":8000")
}

//http://127.0.0.1:8000/user/zhangsan/male
//out:  zhangsan is /male

//http://127.0.0.1:8000/user/zhangsan/male/xyz
//out:  zhangsan is /male
