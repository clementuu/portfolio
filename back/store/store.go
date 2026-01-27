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
