package projet

import (
	"back/model"
	"back/store"
	"testing"
)

func TestMain(m *testing.M) {
	ramStore := store.NewRAMStore()

	ramStore.Projets = []model.Projet{
		{ID: 1, Name: "Maquette 2cv", Sujet: "Monter une maquette de 2cv avec le padre"},
		{ID: 2, Name: "Cheveux longs", Resume: "Laisser pousser les cheveux jusqu'à marcher dessus, me briser le cou et remporter un Darwin awards"},
	}

	Setup(ramStore)

	m.Run()
}
