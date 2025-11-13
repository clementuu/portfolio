package store

import (
	"back/constants"
	"back/model"
	"fmt"
)

const ErrNoProjet = "aucun projet trouvé pour l'id %d"

var projetsData = []model.Projet{
	{ID: 1, Name: "Escarcelle", Image: "", Competences: []model.Competence{GoCompetence, RebolCompetence, SvelteCompetence, GestionProjetCompetence}, Desc: "Escarcelle est une application de gestion d'épiceries solidaire. Ce projet concerne la partie web du logiciel, initalement développée en rebol, le code est progessivement migré vers du go pour le back et du svelte pour le front. L'application repose sur une base de données MySQL."},
	{ID: 2, Name: "Escarcelle Caisse", Image: constants.CaisseLogo, Competences: []model.Competence{GoCompetence}, Desc: "Ce projet est une application de caisse dédiée aux épiceries solidaires, conçue pour répondre aux besoins spécifiques de ces structures à vocation sociale. L'objectif principal est de proposer un outil simple, fiable et adapté à leur fonctionnement, tout en garantissant une expérience utilisateur fluide et intuitive."},
	{ID: 3, Name: "Jeu Anti-Gaspi", Image: constants.AntiGapsiLogo, Competences: []model.Competence{JavaScriptCompetence, HTMLCompetence, CSSCompetence}, Desc: "J'ai développé ce jeu flash dans le cadre de mon service civique. Ma mission était de réaliser une étude du gaspillage alimentaire dans les écoles de la commune de Saint-Jean de Védas, dans l'aglomération de Montpellier. Une fois la partie statistiques achevée il me restait suffisamment de temps pour développer un petit jeu, en javascript/html/css, puis de le présenter aux enfants lors d'atelier de sensibilisation aux éco-gestes."},
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
