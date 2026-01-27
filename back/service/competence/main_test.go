package competence

import (
	"back/model"
	"back/store"
	"testing"
)

func TestMain(m *testing.M) {
	ramStore := store.NewRAMStore()

	ramStore.Competences = []model.Competence{
		{ID: 1, Name: "Go", Rating: 4},
		{ID: 2, Name: "Docker", Rating: 2},
	}

	Setup(ramStore)

	m.Run()
}
