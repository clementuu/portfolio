package service

import "back/model"

type Projet struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	Image       string       `json:"image"`
	Competences []Competence `json:"competences"`
	Desc        string       `json:"desc"`
}

func mapProjet(p model.Projet) Projet {
	var projet Projet
	projet.ID = p.ID
	projet.Name = p.Name
	projet.Image = p.Image
	projet.Competences = getCompetences(p.Competences)
	projet.Desc = p.Desc
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
