package db

import (
	"blogger/model"
	"testing"
	"time"
)

func init() {
	// parseTime = True 将mysql中的时间类开，自动解析为go结构体中的时间类型，否则报错
	dns := "root:root@tcp(127.0.0.1:3306)/blogger?parseTime=true&charset=utf8"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

// 测试插入文章
func TestInsertArticle(t *testing.T) {
	// 构建对象
	artical := &model.ArticleDetail{}
	artical.ArticleInfo.CategoryId = 1
	artical.ArticleInfo.CommentCount = 0
	artical.Content = "一股做起，别来无恙..."
	artical.ArticleInfo.CreateTime = time.Now()
	artical.ArticleInfo.Title = "hai qilai"
	artical.ArticleInfo.UserName = "sun"
	artical.ArticleInfo.ViewCount = 1
	artical.ArticleInfo.Summary = "abc"
	id, err := InsertArticle(artical)
	if err != nil {
		panic(err)
	}
	t.Logf("article_id:%d\n", id)
}

func TestGetArticleList(t *testing.T) {
	articalList, err := GetArticleList(0, 15)
	if err != nil {
		panic(err)
	}
	for _, article := range articalList {
		t.Logf("article:%#v\n", article)
	}
}

func TestGetArticleDetail(t *testing.T) {
	detail, err := GetArticleDetail(1)
	if err != nil {
		panic(err)
	}
	t.Logf("detail:%#v\n", detail)
}

func TestGetAllCategorylist(t *testing.T) {
	articleList, err := GetArticleListByCategoryId(1, 0, 15)
	if err != nil {
		panic(err)
	}
	for _, article := range articleList {
		t.Logf("article:%#v\n", article)
	}
}
