package db

import "testing"

func init() {
	// parseTime = True 将mysql中的时间类开，自动解析为go结构体中的时间类型，否则报错
	dns := "root:root@tcp(127.0.0.1:3306)/blogger?parseTime=true&charset=utf8"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

// 获取单个分类信息
func TestGetCategoryById(t *testing.T) {
	category, err := GetCategoryById(1)
	if err != nil {
		panic(err)
	}
	t.Logf("category:%#v\n", category)
}

func TestGetCategoryList(t *testing.T) {
	var categoryIds []int64
	categoryIds = append(categoryIds, 1, 2, 3)
	list, err := GetCategorylist(categoryIds)
	if err != nil {
		panic(err)
	}
	for idx, category := range list {
		t.Logf("category %d :%#v\n", idx, category)
	}
}

func TestGetAllCategoryList(t *testing.T) {
	list, err := GetAllCategorylist()
	if err != nil {
		panic(err)
	}
	for idx, category := range list {
		t.Logf("category %d :%#v\n", idx, category)
	}
}
