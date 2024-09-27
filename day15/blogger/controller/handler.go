package controller

import (
	"blogger/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 访问主页的控制器
func IndexHandler(c *gin.Context) {
	//从Service取数据
	// 加载了文章数据
	articleRecordList, err := service.GetAllArticleRecordList(0, 15)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	// 加载分类数据
	categoryList, err := service.GetAllCategoryList()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}

	c.HTML(http.StatusOK, "views/index.html", gin.H{
		"article_list":  articleRecordList,
		"category_list": categoryList,
	})

	// gin.H本质上是一个map
	//var data map[string]interface{} = make(map[string]interface{}, 16)
	//data["article_list"] = articleRecordList
	//data["category_list"] = categoryList
	//c.HTML(http.StatusOK, "views/index. html", data)
}

// 点击分类云，进行分类
func CategoryListHandler(c *gin.Context) {
	categoryIdStr := c.Query("category_id")
	if categoryIdStr == "" {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
	}
	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	//根据分类id, 获取文章列表
	articleRecordList, err := service.GetArticleRecordListByCategoryId(categoryId, 0, 15)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}

	// 再次加载分类数据
	categoryList, err := service.GetAllCategoryList()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}

	c.HTML(http.StatusOK, "views/index.html", gin.H{
		"article_list":  articleRecordList,
		"category_list": categoryList,
	})
}
