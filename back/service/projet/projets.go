package projet

import (
	"back/model"
	"back/service/competence"
)

type Model struct {
	ID          int                `json:"id"`
	Name        string             `json:"name"`
	Sujet       string             `json:"sujet"`
	Image       string             `json:"image"`
	Competences []competence.Model `json:"competences"`
	Resume      string             `json:"resume"`
	Template    string             `json:"tmpl"`
}

func mapProjet(p model.Projet) Model {
	var projet Model
	projet.ID = p.ID
	projet.Name = p.Name
	projet.Sujet = p.Sujet
	projet.Image = p.Image
	projet.Competences = competence.ToModels(p.Competences)
	projet.Resume = p.Resume
	projet.Template = p.Template
	return projet
}

func GetAll() (projets []Model) {
	ps := storage.GetProjets()
	for _, p := range ps {
		projet := mapProjet(p)
		projets = append(projets, projet)
	}
	return projets
}

func GetNames() []model.MiniProjet {
	return storage.GetProjetsNames()
}

func Get(id int) (Model, error) {
	p, err := storage.GetProjet(id)
	if err != nil {
		return Model{}, err
	}
	return mapProjet(p), nil
}
