package store

import (
	"back/constants"
	"back/model"
	tmpl "back/tmpl/projets"
	"fmt"
)

const ErrNoProjet = "aucun projet trouvé pour l'id %d"

var (
	projetEscarcelle = model.Projet{ID: 1, Name: "Escarcelle", Sujet: "Gestion d'épicerie solidaire", Image: constants.SiLogo, Competences: []model.Competence{goCompetence, rebolCompetence, webCompetence, svelteCompetence, sqlCompetence, bashCompetence, dockerCompetence, gestionProjetCompetence, flexibiliteCompetence}, Resume: constants.ResumEscarcelle, Template: tmpl.EscarcelleHTML}
	projetCaisse     = model.Projet{ID: 2, Name: "Caisse", Sujet: "Refonte de la caisse Escarcelle", Image: constants.CaisseLogo, Competences: []model.Competence{goCompetence, sqlCompetence, dockerCompetence, creativiteCompetence, flexibiliteCompetence, espritCritiqueCompetence}, Resume: constants.ResumCaisse, Template: tmpl.CaisseHTML}
	projetStats      = model.Projet{ID: 3, Name: "Mission TousAntiGaspi", Sujet: "Analyse statistique du gaspillage alimentaire", Image: constants.SaintJeanLogo, Competences: []model.Competence{rCompetence, pythonCompetence, espritCritiqueCompetence}, Resume: constants.ResumMissionAntiGaspi, Template: tmpl.MissionAntiGaspiHTML}
	projetAntiGaspi  = model.Projet{ID: 4, Name: "Jeu Anti-Gaspi", Sujet: "Jeu éducatif en 2D", Image: constants.AntiGapsiLogo, Competences: []model.Competence{javaScriptCompetence, webCompetence, creativiteCompetence}, Resume: constants.ResumJeuAntiGaspi, Template: tmpl.JeuAntiGaspiHTML}
	projetPMT        = model.Projet{ID: 6, Name: "Project Management Tool", Sujet: "Gestion de projet collaboratif", Image: constants.PMTLogo, Competences: []model.Competence{springCompetence, javaCompetence, angularCompetence, webCompetence, sqlCompetence, dockerCompetence, espritCritiqueCompetence, creativiteCompetence}, Resume: constants.ResumPMT, Template: tmpl.PmtHTML}
)

var projetsList = []model.Projet{
	projetEscarcelle,
	projetCaisse,
	projetStats,
	projetAntiGaspi,
	projetPMT,
}

func (r *RAMStore) GetProjets() []model.Projet {
	return r.Projets
}

func (r *RAMStore) GetProjetsNames() (mps []model.MiniProjet) {
	for _, p := range r.Projets {
		var mp model.MiniProjet
		mp.ID = p.ID
		mp.Name = p.Name
		mps = append(mps, mp)
	}
	return mps
}

func (r *RAMStore) GetProjet(id int) (model.Projet, error) {
	for _, p := range r.Projets {
		if p.ID == id {
			return p, nil
		}
	}
	return model.Projet{}, fmt.Errorf(ErrNoProjet, id)
}
