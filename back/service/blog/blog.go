package blog

import "back/model"

func GetArticles() []model.Article {
	return storage.GetArticles()
}

func GetArticleByID(id int) (model.Article, error) {
	return storage.GetArticleByID(id)
}
