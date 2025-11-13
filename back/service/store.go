package service

import (
	"back/model"
	"back/store"
)

type webStore interface {
	GetCompetences() []model.Competence
	GetProjets() []model.Projet
}

var storage webStore

func Init() {
	storage = store.NewRAMStore()
}
