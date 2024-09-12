package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// API 参数

func main() {
	// 创建一个默认的路由引擎
	// 默认使用了2个中间件Logger(), Recovery,
	r := gin.Default()
	//限制表单上传大小 8M, 默认为32M
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		// 表单取文件
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		//传到项目的根目录，名字就用本身的
		c.SaveUploadedFile(file, "./page/"+file.Filename)

		c.String(http.StatusOK,
			fmt.Sprintf(" file: %s is uploaded. \n", file.Filename))
	})
	r.POST("/uploads", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		}
		// 获取所有图片
		files := form.File["files"]
		// 遍历所有图片
		for _, file := range files {
			if err := c.SaveUploadedFile(file, "./page/"+file.Filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload %s err: %s",
					file.Filename, err.Error()))
				return
			}
		}

		c.String(http.StatusOK,
			fmt.Sprintf("  uploaded ok %d files. \n", len(files)))
	})
	r.Run(":8000")
}

//http://127.0.0.1:8000/welcome?name=Naccy
//out:  Hello, Naccy
