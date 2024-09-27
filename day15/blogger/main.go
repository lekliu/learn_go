package main

import (
	"blogger/controller"
	"blogger/dao/db"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	dns := "root:root@tcp(127.0.0.1:3306)/blogger?parseTime=true&charset=utf8"
	err := db.Init(dns)
	if err != nil {
		panic(err)
	}
	// 加载静态文件
	router.Static("/static", "./static")

	// 加载模板
	router.LoadHTMLGlob("views/*")

	router.GET("/", controller.IndexHandler)
	router.GET("/category/", controller.CategoryListHandler)

	router.Run(":8000")
}
