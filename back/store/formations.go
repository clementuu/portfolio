package store

import "back/model"

var (
	Master  = model.Formation{Intitule: "Master Expert en Ingénierie Logiciel", Etablissement: "ISCOD", Periode: "2023 - 2025"}
	Licence = model.Formation{Intitule: "Licence Mathématique et Informatique appliquées aux SHS", Etablissement: "Paul Valéry Montpellier 3", Periode: "2019 - 2022"}
	Lycee   = model.Formation{Intitule: "Baccalauréat Scientifique", Etablissement: "Lycée Laetitia Bonaparte", Periode: "2013 - 2016"}
)

var formationsList = []model.Formation{Master, Licence, Lycee}

func (s *RAMStore) GetFormations() []model.Formation {
	return formationsList
}
