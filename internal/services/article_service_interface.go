package services

type ArticleServiceInterface interface {

	// LikeArticle 文章点赞
	LikeArticle(articleId int64, userId int64) (err error)
}
