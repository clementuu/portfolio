package service

import (
	"back/model"
	"back/store"
)

type webStore interface {
	GetCompetences() []model.Competence
	GetProjets() []model.Projet
	GetProjetsNames() (mps []model.MiniProjet)
	GetProjet(id int) (model.Projet, error)
	GetFormations() []model.Formation
	GetExperiences() []model.Experience
}

var storage webStore

func Init() {
	storage = store.NewRAMStore()
}
