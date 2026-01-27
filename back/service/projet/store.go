package projet

import (
	"back/model"
)

type webStore interface {
	GetProjets() []model.Projet
	GetProjetsNames() (mps []model.MiniProjet)
	GetProjet(id int) (model.Projet, error)
}

var storage webStore

func Setup(store webStore) {
	storage = store
}
