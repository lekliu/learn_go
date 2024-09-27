package db

import (
	"blogger/model"
	"github.com/jmoiron/sqlx"
)

// 分类相关的操作（添加、查询、查1个分类、查多个分类、查所有分类）

// 添加分类
func InsertCategory(category *model.Category) (categoryId int64, err error) {
	sqlstr := "insert into category(category_name,category_no) values (?,?)"
	result, err := DB.Exec(sqlstr, category.CategoryName, category.Category_no)
	if err != nil {
		return 0, err
	}
	categoryId, err = result.LastInsertId()
	return
}

// 获取单个文章分类
func GetCategoryById(categoryId int64) (category *model.Category, err error) {
	category = &model.Category{}
	sqlstr := "select id,category_name,category_no from category where id = ?"
	err = DB.Get(category, sqlstr, categoryId)
	return
}

// 获取多个文章分类
func GetCategorylist(categoryIds []int64) (categoryList []*model.Category, err error) {
	sqlstr, args, err := sqlx.In("select id,category_name,category_no from category where id in (?)", categoryIds)
	if err != nil {
		return nil, err
	}
	//查询
	err = DB.Select(&categoryList, sqlstr, args...)
	return
}

// 获取所有文章分类
func GetAllCategorylist() (categoryList []*model.Category, err error) {
	sqlstr := "select id,category_name,category_no from category order by category_no asc"
	err = DB.Select(&categoryList, sqlstr)
	return
}
