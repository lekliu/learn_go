package db

import (
	"blogger/model"
	"errors"
)

// 插入文章
func InsertArticle(artical *model.ArticleDetail) (articalId int64, err error) {
	// 加个验证
	if artical == nil {
		return
	}
	sqlstr := `insert into article
    	(content,summary,title,username,category_id,view_count,comment_count)
    	values (?,?,?,?,?,?,?)`
	result, err := DB.Exec(sqlstr, artical.Content, artical.Summary, artical.Title, artical.UserName,
		artical.ArticleInfo.CategoryId, artical.ViewCount, artical.CommentCount)
	if err != nil {
		return 0, err
	}
	articalId, err = result.LastInsertId()
	return
}

// 获取文章列表，做个分页
func GetArticleList(pageNum, pageSize int) (articleList []*model.ArticleInfo, err error) {
	if pageNum < 0 || pageSize <= 0 {
		return
	}
	// 时间降序排序
	sqlstr := `select id,category_id,title,view_count,comment_count,username,summary,create_time
		from article 
		where view_count > 0 and status = 1
		limit ?,?`
	err = DB.Select(&articleList, sqlstr, pageNum, pageSize)
	return
}

// 根据文章ID, 查询单个文章
func GetArticleDetail(articleId int64) (articleDetail *model.ArticleDetail, err error) {
	if articleId < 0 {
		return nil, errors.New("invalid article ID")
	}
	articleDetail = &model.ArticleDetail{}

	// SQL 查询中的列名必须匹配结构体中的 `db` 标签
	sqlstr := `SELECT id, category_id, title, summary, view_count, comment_count, username, create_time, content 
               FROM article 
               WHERE status=1 AND id = ?`

	err = DB.Get(articleDetail, sqlstr, articleId)
	if err != nil {
		return nil, err
	}
	return articleDetail, nil
}

// 根据分类ID，查询这一类文章
func GetArticleListByCategoryId(categoryId int64, pageNum, pageSize int) (articleList []*model.ArticleInfo, err error) {
	if pageNum < 0 || pageSize <= 0 {
		return
	}
	// 时间降序排序
	sqlstr := `select id,category_id,title,view_count,comment_count,username,summary,create_time
		from article 
		where category_id=? and status=1
		limit ?,?`
	err = DB.Select(&articleList, sqlstr, categoryId, pageNum, pageSize)
	return
}
