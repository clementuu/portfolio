package blog

import "back/model"

type blogStore interface {
	GetArticles() []model.Article
	GetArticleByID(id int) (model.Article, error)
}

var storage blogStore

func Setup(store blogStore) {
	storage = store
}
