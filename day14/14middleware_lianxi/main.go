package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// 定义中间件
func timeCount(c *gin.Context) {
	t := time.Now()
	//执行函数
	c.Next()
	t2 := time.Since(t)
	fmt.Println("程序用时:", t2.String())
}

func main() {
	// 创建一个默认的路由引擎
	// 默认使用了2个中间件Logger(), Recovery,
	r := gin.Default()
	// 注册中间件
	r.Use(timeCount)
	//{} 为了代码规范
	shoppingGroup := r.Group("/shopping")
	{
		shoppingGroup.GET("/index", handler1)

		shoppingGroup.GET("/home", handler2)
	}

	r.Run(":8000")
}

func handler1(c *gin.Context) {
	time.Sleep(time.Second)
	c.JSON(200, gin.H{"测试": "中间件"})
}

func handler2(c *gin.Context) {
	time.Sleep(time.Second)
	c.JSON(200, gin.H{"测试": "中间件"})
}

// http://127.0.0.1:8000/shopping/index
// http://127.0.0.1:8000/shopping/home
