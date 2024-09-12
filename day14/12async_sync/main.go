package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	// 创建一个默认的路由引擎
	// 默认使用了2个中间件Logger(), Recovery,
	r := gin.Default()
	// 1. 异步
	r.GET("/long_async", func(c *gin.Context) {
		//需要一个副本
		copyContext := c.Copy()
		go func() {
			time.Sleep(3 * time.Second)
			log.Println("异步执行：" + copyContext.Request.URL.Path)
		}()
	})

	//同步
	r.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		log.Println("同步执行：" + c.Request.URL.Path)
		//c.JSON(http.StatusOK, gin.H{"message": "someJSON", "status": 200})
	})

	r.Run(":8000")
}

//http://127.0.0.1:8000/long_async
//http://127.0.0.1:8000/long_sync
