package web

import (
	"back/model"
	"back/service/competence"
	"back/service/cv"
	"back/service/projet"
	"back/store"
	"testing"
)

func TestMain(m *testing.M) {
	ramStore := store.NewRAMStore()

	ramStore.Projets = []model.Projet{
		{ID: 1, Name: "Maquette 2cv", Sujet: "Monter une maquette de 2cv avec le padre"},
		{ID: 2, Name: "Cheveux longs", Resume: "Laisser pousser les cheveux jusqu'à marcher dessus, me briser le cou et remporter un Darwin awards"},
	}

	ramStore.Competences = []model.Competence{
		{ID: 1, Name: "Go", Rating: 4},
		{ID: 2, Name: "Docker", Rating: 2},
	}

	ramStore.Formations = []model.Formation{
		{Intitule: "Master", Periode: "2023 - 2025"},
		{Intitule: "Licence", Periode: "2019 - 2022"},
		{Intitule: "Lycee", Etablissement: "Bonaparte", Periode: "2013 - 2016"},
	}

	ramStore.Experiences = []model.Experience{
		{Intitule: "Pizzaiolo", Type: "Acrobatique"},
		{Intitule: "Rockstar", Taches: []string{"Fumer un max de clopes", "Crier très fort dans un micro", "Faire coucou aux fans"}},
	}

	competence.Setup(ramStore)
	projet.Setup(ramStore)
	cv.Setup(ramStore)

	m.Run()
}
