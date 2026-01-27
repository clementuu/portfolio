package competence

import "back/model"

type webStore interface {
	GetCompetence(id int) (model.Competence, error)
	GetCompetences() []model.Competence
}

var storage webStore

func Setup(store webStore) {
	storage = store
}
