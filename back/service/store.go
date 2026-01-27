package service

import (
	"back/model"
)

type webStore interface {
	GetCompetence(id int) (model.Competence, error)
	GetCompetences() []model.Competence
	GetProjets() []model.Projet
	GetProjetsNames() (mps []model.MiniProjet)
	GetProjet(id int) (model.Projet, error)
	GetFormations() []model.Formation
	GetExperiences() []model.Experience
}

var storage webStore

func Init(store webStore) {
	storage = store
}
