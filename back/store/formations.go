package store

import "back/model"

var (
	master  = model.Formation{Intitule: "Master Expert en Ingénierie Logiciel", Etablissement: "ISCOD", Periode: "2023 - 2025"}
	licence = model.Formation{Intitule: "Licence Mathématique et Informatique appliquées aux SHS", Etablissement: "Paul Valéry Montpellier 3", Periode: "2019 - 2022"}
	lycee   = model.Formation{Intitule: "Baccalauréat Scientifique", Etablissement: "Lycée Laetitia Bonaparte", Periode: "2013 - 2016"}
)

var formationsList = []model.Formation{master, licence, lycee}

func (r *RAMStore) GetFormations() []model.Formation {
	return r.Formations
}
