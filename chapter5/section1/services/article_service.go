package services

import "github.com/yourname/reponame/models"

// PostArticleHandlerで使うことを想定したサービス
// 引数の情報をもとに新しい記事を作り、結果を返却
func PostArticleService(article models.Article) (models.Article, error) {
	// TODO: 実装
	return models.Article{}, nil
}

// ArticleListHandlerで使うことを想定したサービス
// 指定pageの記事一覧を返却
func GetArticleListService(page int) ([]models.Article, error) {
	// TODO: 実装
	return nil, nil
}

// ArticleDetailHandlerで使うことを想定したサービス
// 指定IDの記事情報を返却
func GetArticleService(articleID int) (models.Article, error) {
	// TODO: 実装
	return models.Article{}, nil
}

// PostNiceHandlerで使うことを想定したサービス
// 指定IDの記事のいいね数を+1して、結果を返却
func PostNiceService(article models.Article) (models.Article, error) {
	// TODO: 実装
	return models.Article{}, nil
}
