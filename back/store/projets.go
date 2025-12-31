package store

import (
	"back/constants"
	"back/model"
	"back/tmpl"
	"fmt"
)

const ErrNoProjet = "aucun projet trouvé pour l'id %d"

var (
	ProjetEscarcelle = model.Projet{ID: 1, Name: "Escarcelle", Sujet: "Gestion d'épicerie solidaire", Image: "", Competences: []model.Competence{GoCompetence, RebolCompetence, SvelteCompetence, SQLCompetence, GestionProjetCompetence}, Resume: constants.ResumEscarcelle, Template: tmpl.EscarcelleProject}
	ProjetCaisse     = model.Projet{ID: 2, Name: "Caisse", Sujet: "Refonte de la caisse Escarcelle", Image: constants.CaisseLogo, Competences: []model.Competence{GoCompetence, SQLCompetence}, Resume: constants.ResumCaisse, Template: tmpl.CaisseProject}
	ProjetStats      = model.Projet{ID: 3, Name: "Mission Anti-Gaspi", Sujet: "Analyse statistique du gaspillage alimentaire", Image: constants.SaintJeanLogo, Competences: []model.Competence{RCompetence, PythonCompetence}, Resume: constants.ResumMissionAntiGaspi, Template: tmpl.MissionAntiGaspiHTML}
	ProjetAntiGaspi  = model.Projet{ID: 4, Name: "Jeu Anti-Gaspi", Sujet: "Jeu éducatif en 2D", Image: constants.AntiGapsiLogo, Competences: []model.Competence{JavaScriptCompetence, HTMLCompetence, CSSCompetence}, Resume: constants.ResumJeuAntiGaspi, Template: tmpl.JeuAntiGaspiHTML}
	ProjetRGC        = model.Projet{ID: 5, Name: "RGC", Image: "", Competences: []model.Competence{AngularCompetence, SQLCompetence, GoCompetence}, Resume: constants.ResumRGC, Template: tmpl.RGCHTML}
	ProjetPMT        = model.Projet{ID: 6, Name: "Project Management Tool", Image: "", Competences: []model.Competence{SpringCompetence, JavaCompetence, AngularCompetence, SQLCompetence}, Resume: constants.ResumPMT, Template: tmpl.PMTHTML}
)

var projetsData = []model.Projet{
	ProjetEscarcelle,
	ProjetCaisse,
	ProjetStats,
	ProjetAntiGaspi,
	ProjetRGC,
	ProjetPMT,
}

func (s *RAMStore) GetProjets() []model.Projet {
	return projetsData
}

func (s *RAMStore) GetProjetsNames() (mps []model.MiniProjet) {
	for _, p := range projetsData {
		var mp model.MiniProjet
		mp.ID = p.ID
		mp.Name = p.Name
		mps = append(mps, mp)
	}
	return mps
}

func (s *RAMStore) GetProjet(id int) (model.Projet, error) {
	for _, p := range projetsData {
		if p.ID == id {
			return p, nil
		}
	}
	return model.Projet{}, fmt.Errorf(ErrNoProjet, id)
}
