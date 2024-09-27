package model

import (
	"time"
)

//    id  category_id  content  title   view_count  comment_count  username  status  summary  create_time  update_time
//------  -----------  -------  ------  ----------  -------------  --------  ------  -------  -----------  -------------

// 定义文章的结构体
type ArticleInfo struct {
	Id           int64     `db:"id"`
	CategoryId   int64     `db:"category_id"`
	Title        string    `db:"title"`
	Summary      string    `db:"summary"`
	ViewCount    uint32    `db:"view_count"`
	CommentCount uint32    `db:"comment_count"`
	UserName     string    `db:"username"`
	CreateTime   time.Time `db:"create_time"`
}

// 用于文章详情页的实体
// 为了提升效率，将Content单独拿出来
type ArticleDetail struct {
	ArticleInfo
	//文章内容
	Content string `db:"content"`
	//Category
}

// 用于文章上下页
type ArticleRecord struct {
	ArticleInfo
	Category
}
