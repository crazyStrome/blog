package service

import (
	"blog/model"
	"blog/dao"
)
// MapAuthorIDAndArticleID 把作者id和文章id映射到数据库中
func MapAuthorIDAndArticleID(authorid, articleid int64) int64 {
	return dao.AddMapAuthorIDAndArticleID(authorid, articleid)
}
// GetArticleIDListByAuthorID 获取作者的文章id列表
func GetArticleIDListByAuthorID(authorid int64) []int64 {
	return dao.QueryArticleIDsByAuthorID(authorid)
}
// GetArticleByID 获取文章通过文章id
func GetArticleByID(articleid int64) (article model.Article, err error) {
	return dao.QueryArticleByID(articleid)
}
// AddArticle 添加文章
func AddArticle(article model.Article) int64 {
	return dao.AddArticle(article)
}