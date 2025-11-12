package service

import (
	"back/model"
	"back/store"
)

type webStore interface {
	GetCompetences() []model.Competence
}

var storage webStore

func Init() {
	storage = store.NewRAMStore()
}
