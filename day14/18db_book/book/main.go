package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 初始化数据库
	err := initDB()
	if err != nil {
		panic(err)
	}
	// 创建一个默认的路由引擎
	r := gin.Default()
	//加载页面
	r.LoadHTMLGlob("./book/templates/*")
	//查询所有图书
	r.GET("/book/list", bookListHandler)
	r.Run(":8000")
}

func bookListHandler(c *gin.Context) {
	bookList, err := queryAllBook()
	//fmt.Printf("%#v\n", bookList)
	//for _, book := range bookList {
	//	fmt.Printf("%#v\n", book)
	//}
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err,
		})
		return
	}
	c.HTML(http.StatusOK, "book_list.html", gin.H{
		"code": 0,
		"data": bookList,
	})
}

// http://loacal:8000/book/list
