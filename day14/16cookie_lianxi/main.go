package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 定义中间件
func CheckAuthority(c *gin.Context) {
	if cookie, err := c.Cookie("login_cookie"); err == nil {
		if cookie == "123456" {
			//执行后续函数
			c.Next()
			return
		}
	}
	c.JSON(http.StatusUnauthorized, gin.H{
		"error": "StatusUnauthorized",
	})
	c.Abort()
	return
}

func main() {
	// 创建一个默认的路由引擎
	// 默认使用了2个中间件Logger(), Recovery,
	r := gin.Default()

	r.GET("/login", func(c *gin.Context) {
		// 获取客户端是否携带cookie
		_, err := c.Cookie("login_cookie")
		if err != nil {
			c.SetCookie("login_cookie", "123456", 60, "/",
				"localhost", false, true)
		}
		c.String(http.StatusOK, "Login successful")
	})

	r.GET("/home", CheckAuthority, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "home",
		})
	})
	r.Run(":8000")
}

// http://localhost:8000/home
// http://localhost:8000/login
// http://localhost:8000/home
