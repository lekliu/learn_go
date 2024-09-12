package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 定义接收数据的结构体
type Login struct {
	// binding:"required" 修饰的字段，是必选字段，若接收为空值，则报错
	User     string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main() {
	// 创建一个默认的路由引擎
	// 默认使用了2个中间件Logger(), Recovery,
	r := gin.Default()
	// JSON 绑定
	r.GET("/loginUri/:user/:password", func(c *gin.Context) {
		// 声明接收的变量
		var uri Login
		//将request中的body中的数据，自动按照json格式解析到结构体
		if err := c.ShouldBindUri(&uri); err != nil {
			// 返回错误信息
			// gin.H 封装了生成json的数扰的工具
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//判断用户名与密码是否正确
		if uri.User != "root" || uri.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": 304})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": 200})
	})

	r.Run(":8000")
}

// C:\>curl http://127.0.0.1:8000/loginUri/root/admin
// {"status":200}
