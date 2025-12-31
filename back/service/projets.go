package service

import "back/model"

type Projet struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	Sujet       string       `json:"sujet"`
	Image       string       `json:"image"`
	Competences []Competence `json:"competences"`
	Resume      string       `json:"resume"`
	Template    string       `json:"tmpl"`
}

func mapProjet(p model.Projet) Projet {
	var projet Projet
	projet.ID = p.ID
	projet.Name = p.Name
	projet.Sujet = p.Sujet
	projet.Image = p.Image
	projet.Competences = getCompetences(p.Competences)
	projet.Resume = p.Resume
	projet.Template = p.Template
	return projet
}

func GetAllProjets() (projets []Projet) {
	ps := storage.GetProjets()
	for _, p := range ps {
		projet := mapProjet(p)
		projets = append(projets, projet)
	}
	return projets
}

func GetProjetsNames() []model.MiniProjet {
	return storage.GetProjetsNames()
}

func GetProjet(id int) (Projet, error) {
	p, err := storage.GetProjet(id)
	if err != nil {
		return Projet{}, err
	}
	return mapProjet(p), nil
}
