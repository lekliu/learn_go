package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// 定义中间件
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了...")
		// 设置变量到Context的key中，可以通过Get取
		c.Set("request", "中间件")
		//执行函数
		c.Next()
		//中间件执行完毕后续的一些事
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕,", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2.String())
	}
}

func main() {
	// 创建一个默认的路由引擎
	// 默认使用了2个中间件Logger(), Recovery,
	r := gin.Default()
	// 注册中间件
	r.Use(Middleware())
	//{} 为了代码规范
	{
		r.GET("/middleware", func(c *gin.Context) {
			//取值
			req, _ := c.Get("request")
			fmt.Println("request:", req)
			c.JSON(200, gin.H{"request": req})
		})

		//在GET方法中，定义局部中间件
		r.GET("/middleware2", Middleware(), func(c *gin.Context) {
			//取值
			req, _ := c.Get("request")
			fmt.Println("request:", req)
			c.JSON(200, gin.H{"request": req})
		})
	}

	r.Run(":8000")
}

// http://127.0.0.1:8000/middleware
// {"request":"中间件"}
