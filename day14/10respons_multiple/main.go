package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

func main() {
	// 创建一个默认的路由引擎
	// 默认使用了2个中间件Logger(), Recovery,
	r := gin.Default()
	// 1. JSON
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "someJSON", "status": 200})
	})

	// 2. 结构体响应
	r.GET("/someStruct", func(c *gin.Context) {
		var msg struct {
			Name    string
			Message string
			Number  int
		}
		msg.Name = "root"
		msg.Message = "message"
		msg.Number = 123
		c.JSON(http.StatusOK, msg)
	})

	// 3.XML
	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "someXML", "status": 200})
	})

	// 4. YAML响应
	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"name": "zhangsan", "age": 200})
	})

	// 5.protobuf 格式，谷歌开发的高效存储读取的工具
	r.GET("/someProtobuf", func(c *gin.Context) {
		response := []int64{int64(1), int64(2), int64(3), int64(4), int64(5), int64(6)}
		//定义数据
		label := "label"
		data := &protoexample.Test{
			Label: &label,
			Reps:  response,
		}
		c.ProtoBuf(http.StatusOK, data)
	})
	r.Run(":8000")
}

// http://127.0.0.1:8000/someJSON
// {"message":"someJSON","status":200}

//http://127.0.0.1:8000/someStruct
//http://127.0.0.1:8000/someXML
//http://127.0.0.1:8000/someYAML
//http://127.0.0.1:8000/someProtobuf
