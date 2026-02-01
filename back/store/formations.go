package store

import (
	"back/constants"
	"back/model"
	tmpl "back/tmpl/etablissements"
)

var (
	master  = model.Formation{ID: 1, Intitule: "Master Expert en Ingénierie Logiciel", Etablissement: "ISCOD", Periode: "2023 - 2025", Logo: constants.ISCODLogo, Presentation: tmpl.IscodPresentation}
	licence = model.Formation{ID: 2, Intitule: "Licence Mathématique et Informatique appliquées aux SHS", Etablissement: "Paul Valéry Montpellier 3", Periode: "2019 - 2022", Logo: constants.PaulVaLogo, Presentation: tmpl.PaulVaPresentation}
	lycee   = model.Formation{ID: 3, Intitule: "Baccalauréat Scientifique", Etablissement: "Lycée Laetitia Bonaparte", Periode: "2013 - 2016", Logo: constants.LaetitiaLogo, Presentation: tmpl.LaetitiaPresentation}
)

var formationsList = []model.Formation{master, licence, lycee}

func (r *RAMStore) GetFormations() []model.Formation {
	return r.Formations
}
