package service

import (
	"blogger/dao/db"
	"blogger/model"
)

// 获取所有分类
func GetAllCategoryList() (categoryList []*model.Category, err error) {
	categoryList, err = db.GetAllCategorylist()
	if err != nil {
		return nil, err
	}
	return categoryList, nil
}
