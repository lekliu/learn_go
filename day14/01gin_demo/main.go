package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建一个默认的路由引擎
	// 默认使用了2个中间件Logger(), Recovery, 如果不用中间件，可执行 gin.New()
	r := gin.Default()
	// GET：请求方式；/hello：请求的路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	r.GET("/hello", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world!")
	})
	r.POST("xxxPost", getting)
	r.PUT("xxxPUT", func(c *gin.Context) {
		//TODO
	})
	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	// r.Run()
	r.Run(":8000")
}

func getting(c *gin.Context) {

}
