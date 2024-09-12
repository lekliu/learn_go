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
	r.POST("/form", func(c *gin.Context) {
		//表单参数设置默认值
		type1 := c.DefaultPostForm("type", "alert")
		// 接收其它的
		username := c.PostForm("username")
		password := c.PostForm("password")
		//多选框
		hobbies := c.PostFormArray("hobby")

		c.String(http.StatusOK,
			fmt.Sprintf("type is %s, username is %s, password is %s, hobbys is %v\n",
				type1, username, password, hobbies))
	})
	r.Run(":8000")
}

//http://127.0.0.1:8000/welcome?name=Naccy
//out:  Hello, Naccy
