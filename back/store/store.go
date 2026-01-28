package store

import "back/model"

type RAMStore struct {
	Competences []model.Competence
	Projets     []model.Projet
	Formations  []model.Formation
	Experiences []model.Experience
}

// NewRAMStore crée une nouvelle instance de RAMStore.
func NewRAMStore() *RAMStore {
	return &RAMStore{
		Competences: competencesList,
		Projets:     projetsList,
		Formations:  formationsList,
		Experiences: experiencesList,
	}
}

func SetupTestData() *RAMStore {
	return &RAMStore{
		Projets: []model.Projet{
			{ID: 1, Name: "Maquette 2cv", Sujet: "Monter une maquette de 2cv avec le padre"},
			{ID: 2, Name: "Cheveux longs", Resume: "Laisser pousser les cheveux jusqu'à marcher dessus, me briser le cou et remporter un Darwin awards"},
		},
		Competences: []model.Competence{
			{ID: 1, Name: "Go", Rating: 4},
			{ID: 2, Name: "Docker", Rating: 2},
		},
		Formations: []model.Formation{
			{Intitule: "Master", Periode: "2023 - 2025"},
			{Intitule: "Licence", Periode: "2019 - 2022"},
			{Intitule: "Lycee", Etablissement: "Bonaparte", Periode: "2013 - 2016"},
		},
		Experiences: []model.Experience{
			{Intitule: "Pizzaiolo", Type: "Acrobatique"},
			{Intitule: "Rockstar", Taches: []string{"Fumer un max de clopes", "Crier très fort dans un micro", "Faire coucou aux fans"}},
		},
	}
}
