package store

import (
	"back/constants"
	"back/model"
	tmpl "back/tmpl/etablissements"
)

var (
	master  = model.Formation{ID: 1, Intitule: "Master Expert en Ingénierie Logiciel", Etablissement: "ISCOD", Periode: "2024 - 2026", Logo: constants.ISCODLogo, Presentation: tmpl.IscodPresentation, URL: "https://www.iscod.fr/"}
	licence = model.Formation{ID: 2, Intitule: "Licence Mathématique et Informatique appliquées aux SHS", Etablissement: "Paul Valéry Montpellier 3", Periode: "2019 - 2022", Logo: constants.PaulVaLogo, Presentation: tmpl.PaulVaPresentation, URL: "https://www.univ-montp3.fr"}
	lycee   = model.Formation{ID: 3, Intitule: "Baccalauréat Scientifique", Etablissement: "Lycée Laetitia Bonaparte", Periode: "2013 - 2016", Logo: constants.LaetitiaLogo, Presentation: tmpl.LaetitiaPresentation, URL: "https://llb.ac-corse.fr/"}
)

var formationsList = []model.Formation{master, licence, lycee}

func (r *RAMStore) GetFormations() []model.Formation {
	return r.Formations
}
