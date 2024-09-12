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
	// 加载模板文件
	r.LoadHTMLGlob("templates/*")
	//r.LoadHTMLFiles("templates/index.tmpl")
	r.GET("/index", func(c *gin.Context) {
		//根据文件名渲染
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "我的标题",
		})
	})

	r.Run(":8000")
}

//http://127.0.0.1:8000/index
