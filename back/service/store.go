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
}

var storage webStore

func Init() {
	storage = store.NewRAMStore()
}
