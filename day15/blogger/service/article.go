package service

import (
	"blogger/dao/db"
	"blogger/model"
)

// 获取文章和和他们对应的分类信息
func GetAllArticleRecordList(pageNum, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {
	//1 获取文章列表
	articleInfoList, err := db.GetArticleList(pageNum, pageSize)
	if err != nil {
		return nil, err
	}
	articleRecordList, records, err2, done := GetArticleRecordList(articleInfoList, err, articleRecordList)
	if done {
		return records, err2
	}
	return
}

// 根据分类id,获取该类文章和她们对应的分类信息
func GetArticleRecordListByCategoryId(categoryId int64, pageNum, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {
	//1 获取文章列表
	articleInfoList, err := db.GetArticleListByCategoryId(categoryId, pageNum, pageSize)
	if err != nil {
		return nil, err
	}
	articleRecordList, records, err2, done := GetArticleRecordList(articleInfoList, err, articleRecordList)
	if done {
		return records, err2
	}
	return
}

// 公共函数
func GetArticleRecordList(articleInfoList []*model.ArticleInfo, err error, articleRecordList []*model.ArticleRecord) ([]*model.ArticleRecord, []*model.ArticleRecord, error, bool) {
	if len(articleInfoList) <= 0 {
		return nil, nil, nil, true
	}
	//2 获取文章对应的分类（多个）
	CategoryIds := getCategoryIds(articleInfoList)
	categoryList, err := db.GetCategorylist(CategoryIds)
	if err != nil {
		return nil, nil, err, true
	}
	//将categoryList，转化成map形式
	categoryListMap := make(map[int64]*model.Category)
	for _, categoryInfo := range categoryList {
		categoryListMap[categoryInfo.CategoryId] = categoryInfo
	}
	//3.返回页面，做聚合
	for _, articleInfo := range articleInfoList {
		// 根据档前文章，生成结构体
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *articleInfo,
			Category:    *categoryListMap[articleInfo.CategoryId],
		}
		articleRecordList = append(articleRecordList, articleRecord)
	}
	return articleRecordList, nil, nil, false
}

// 根据多个文章的id, 获取多个分类id的集合
func getCategoryIds(articleInfoList []*model.ArticleInfo) (categoryIds []int64) {
	// 遍历文章，得到第个文章id，放入map中
	categorySet := make(map[int64]struct{})
	for _, article := range articleInfoList {
		categorySet[article.CategoryId] = struct{}{} // 将 CategoryId 作为键插入 map，struct{}{} 是占位符
	}

	for categoryId := range categorySet {
		categoryIds = append(categoryIds, categoryId) // 将去重后的键添加到 slice 中
	}
	return
}
