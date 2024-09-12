package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	// 创建一个默认的路由引擎
	// 默认使用了2个中间件Logger(), Recovery,
	r := gin.Default()
	// 服务端要给客户端Cookie
	r.GET("cookie", func(c *gin.Context) {
		// 获取客户端是否携带cookie
		cookie, err := c.Cookie("key_cookie")
		if err != nil {
			cookie = "NoSet"
			// 设置Cookie
			// maxAge int，单位为秒
			// path,cookie 所在目录
			// domain string 域名
			// secure 是否只能通过https访问
			// httpOnly bool,是否允许别人通过js获取自己的cookie
			c.SetCookie("key_cookie", "value_cookie", 60, "/",
				"localhost", false, true)
		}
		fmt.Println("cookie的值是 ", cookie)
		c.JSON(200, gin.H{"cookie:": cookie})
	})
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

// http://localhost:8000/cookie
