package service

import "back/model"

type Competence struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Image  string `json:"image"`
	Rating int    `json:"rating"`
	Desc   string `json:"desc"`
	Exp    string `json:"exp"`
	Type   string `json:"type"`
}

func getCompetences(cs []model.Competence) (competences []Competence) {
	for _, c := range cs {
		var competence Competence
		competence.ID = c.ID
		competence.Name = c.Name
		competence.Image = c.Image
		competence.Rating = c.Rating
		competence.Desc = c.Template
		competence.Exp = c.Exp
		competence.Type = c.Type.String()
		competences = append(competences, competence)
	}
	return competences
}

func GetAllCompetences() (competences []Competence) {
	cs := storage.GetCompetences()
	return getCompetences(cs)
}
