package store

import (
	"back/constants"
	"back/model"
	tmpl "back/tmpl/projets"
	"fmt"
)

const ErrNoProjet = "aucun projet trouvé pour l'id %d"

var (
	ProjetEscarcelle = model.Projet{ID: 1, Name: "Escarcelle", Sujet: "Gestion d'épicerie solidaire", Image: constants.SiLogo, Competences: []model.Competence{GoCompetence, RebolCompetence, WebCompetence, SvelteCompetence, SQLCompetence, BashCompetence, DockerCompetence, GestionProjetCompetence, FlexibiliteCompetence}, Resume: constants.ResumEscarcelle, Template: tmpl.EscarcelleHTML}
	ProjetCaisse     = model.Projet{ID: 2, Name: "Caisse", Sujet: "Refonte de la caisse Escarcelle", Image: constants.CaisseLogo, Competences: []model.Competence{GoCompetence, SQLCompetence, DockerCompetence, CreativiteCompetence, FlexibiliteCompetence, EspritCritiqueCompetence}, Resume: constants.ResumCaisse, Template: tmpl.CaisseHTML}
	ProjetStats      = model.Projet{ID: 3, Name: "Mission TousAntiGaspi", Sujet: "Analyse statistique du gaspillage alimentaire", Image: constants.SaintJeanLogo, Competences: []model.Competence{RCompetence, PythonCompetence, EspritCritiqueCompetence}, Resume: constants.ResumMissionAntiGaspi, Template: tmpl.MissionAntiGaspiHTML}
	ProjetAntiGaspi  = model.Projet{ID: 4, Name: "Jeu Anti-Gaspi", Sujet: "Jeu éducatif en 2D", Image: constants.AntiGapsiLogo, Competences: []model.Competence{JavaScriptCompetence, WebCompetence, CreativiteCompetence}, Resume: constants.ResumJeuAntiGaspi, Template: tmpl.JeuAntiGaspiHTML}
	ProjetPMT        = model.Projet{ID: 6, Name: "Project Management Tool", Sujet: "Gestion de projet collaboratif", Image: "", Competences: []model.Competence{SpringCompetence, JavaCompetence, AngularCompetence, WebCompetence, SQLCompetence, DockerCompetence, EspritCritiqueCompetence, CreativiteCompetence}, Resume: constants.ResumPMT, Template: tmpl.PmtHTML}
)

var projetsList = []model.Projet{
	ProjetEscarcelle,
	ProjetCaisse,
	ProjetStats,
	ProjetAntiGaspi,
	ProjetPMT,
}

func (s *RAMStore) GetProjets() []model.Projet {
	return projetsList
}

func (s *RAMStore) GetProjetsNames() (mps []model.MiniProjet) {
	for _, p := range projetsList {
		var mp model.MiniProjet
		mp.ID = p.ID
		mp.Name = p.Name
		mps = append(mps, mp)
	}
	return mps
}

func (s *RAMStore) GetProjet(id int) (model.Projet, error) {
	for _, p := range projetsList {
		if p.ID == id {
			return p, nil
		}
	}
	return model.Projet{}, fmt.Errorf(ErrNoProjet, id)
}
