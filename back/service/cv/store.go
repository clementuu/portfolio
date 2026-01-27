package cv

import (
	"back/model"
)

type webStore interface {
	GetFormations() []model.Formation
	GetExperiences() []model.Experience
}

var storage webStore

func Setup(store webStore) {
	storage = store
}
