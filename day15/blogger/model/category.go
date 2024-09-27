package model

// `id``category_name``category_no``create_time``update_time`

// 定义分类结构体
type Category struct {
	CategoryId   int64  `db:"id"`
	CategoryName string `db:"category_name"`
	Category_no  string `db:"category_no"`
}
