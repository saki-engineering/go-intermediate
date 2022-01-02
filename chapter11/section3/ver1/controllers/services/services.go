package services

import "github.com/yourname/reponame/models"

// /article関連を引き受けるサービス
type ArticleServicer interface {
	PostArticleService(article models.Article) (models.Article, error)
	GetArticleListService(page int) ([]models.Article, error)
	GetArticleService(articleID int) (models.Article, error)
	PostNiceService(article models.Article) (models.Article, error)
}

// /commentを引き受けるサービス
type CommentServicer interface {
	PostCommentService(comment models.Comment) (models.Comment, error)
}
