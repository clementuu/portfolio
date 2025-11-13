package store

import (
	"back/constants"
	"back/model"
	"fmt"
)

const ErrNoProjet = "aucun projet trouvé pour l'id %d"

var (
	ProjetEscarcelle = model.Projet{ID: 1, Name: "Escarcelle", Image: "", Competences: []model.Competence{GoCompetence, RebolCompetence, SvelteCompetence, GestionProjetCompetence}, Desc: "Escarcelle est une application de gestion d'épiceries solidaire. Ce projet concerne la partie web du logiciel, initalement développée en rebol, le code est progessivement migré vers du go pour le back et du svelte pour le front. L'application repose sur une base de données MySQL."}
	ProjetCaisse     = model.Projet{ID: 2, Name: "Escarcelle Caisse", Image: constants.CaisseLogo, Competences: []model.Competence{GoCompetence}, Desc: "Ce projet est une application de caisse dédiée aux épiceries solidaires, conçue pour répondre aux besoins spécifiques de ces structures à vocation sociale. L'objectif principal est de proposer un outil simple, fiable et adapté à leur fonctionnement, tout en garantissant une expérience utilisateur fluide et intuitive."}
	ProjetStats      = model.Projet{ID: 3, Name: "Mission Anti-Gaspi", Image: constants.SaintJeanLogo, Competences: []model.Competence{RCompetence, PythonCompetence}, Desc: "Réalisation d'une étude statistique du gaspillage alimentaire dans les cantines scolaires dans la commune de Saint-Jean de Védas, dans l'aglomération de Montpellier. Récolte et analyse de données avec pour outils R et python."}
	ProjetAntiGaspi  = model.Projet{ID: 4, Name: "Jeu Anti-Gaspi", Image: constants.AntiGapsiLogo, Competences: []model.Competence{JavaScriptCompetence, HTMLCompetence, CSSCompetence}, Desc: "J'ai développé ce jeu flash dans le cadre de mon service civique. Une fois la partie statistiques achevée il me restait suffisamment de temps pour développer un petit jeu, en javascript/html/css, puis de le présenter aux enfants lors d'atelier de sensibilisation aux éco-gestes."}
	ProjetRGC        = model.Projet{ID: 5, Name: "RGC", Image: "", Competences: []model.Competence{AngularCompetence, SQLCompetence, GoCompetence}, Desc: "Application de gestion web, développé pour une société de recouvrement de fond. J'ai participé à la maintenance et à quelques évolutions."}
	ProjetPMT        = model.Projet{ID: 6, Name: "Project Management Tool", Image: "", Competences: []model.Competence{SpringCompetence, JavaCompetence, AngularCompetence, SQLCompetence}, Desc: "Projet de création d'une application de gestion de projet, permettant de travailler en équipe, d'organiser ses projets, créer des tâches, les assigners aux participants..."}
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
