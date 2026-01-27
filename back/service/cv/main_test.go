package cv

import (
	"back/model"
	"back/store"
	"testing"
)

func TestMain(m *testing.M) {
	ramStore := store.NewRAMStore()

	ramStore.Formations = []model.Formation{
		{Intitule: "Master", Periode: "2023 - 2025"},
		{Intitule: "Licence", Periode: "2019 - 2022"},
		{Intitule: "Lycee", Etablissement: "Bonaparte", Periode: "2013 - 2016"},
	}
	ramStore.Experiences = []model.Experience{
		{Intitule: "Pizzaiolo", Type: "Acrobatique"},
		{Intitule: "Rockstar", Taches: []string{"Fumer un max de clopes", "Crier très fort dans un micro", "Faire coucou aux fans"}},
	}

	Setup(ramStore)

	m.Run()
}
