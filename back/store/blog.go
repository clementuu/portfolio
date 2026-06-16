package store

import (
	"back/model"
	tmpl "back/tmpl/blog"
	"bytes"
	"errors"
	"html/template"
	"log"
)

var ErrorArticleNotFound = errors.New("Article introuvable")

var (
	premierArticle = model.Article{ID: 1, Titre: "Premier article", Date: "16/06/2026", Description: "Dans cet article, je partage ma vision du développement logiciel et pourquoi j'ai décidé de créer ce portfolio.", Template: tmpl.PremierArticle}
)

var articlesList = []model.Article{
	parseTemplate(premierArticle),
}

func parseTemplate(article model.Article) model.Article {
	t, err := template.New("article").Parse(article.Template)
	if err != nil {
		log.Printf("Erreur parsing template article %d: %v", article.ID, err)
		return article
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, article); err != nil {
		log.Printf("Erreur execution template article %d: %v", article.ID, err)
		return article
	}

	article.Template = tpl.String()

	return article
}

func (r *RAMStore) GetArticles() []model.Article {
	return r.Blog
}

func (r *RAMStore) GetArticleByID(id int) (model.Article, error) {
	for _, article := range r.Blog {
		if article.ID == id {
			return article, nil
		}
	}
	return model.Article{}, ErrorArticleNotFound
}
