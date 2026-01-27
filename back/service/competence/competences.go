package competence

import "back/model"

type Model struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Image    string `json:"image"`
	Rating   int    `json:"rating"`
	Template string `json:"template"`
	Exp      string `json:"exp"`
	Type     string `json:"type"`
}

func ToModels(cs []model.Competence) (competences []Model) {
	for _, c := range cs {
		competence := toModel(c)
		competences = append(competences, competence)
	}
	return competences
}

func toModel(c model.Competence) (competence Model) {
	competence.ID = c.ID
	competence.Name = c.Name
	competence.Image = c.Image
	competence.Rating = c.Rating
	competence.Template = c.Template
	competence.Exp = c.Exp
	competence.Type = c.Type.String()

	return competence
}

func GetAll() (competences []Model) {
	cs := storage.GetCompetences()
	return ToModels(cs)
}

func Get(id int) (Model, error) {
	c, err := storage.GetCompetence(id)
	if err != nil {
		return Model{}, err
	}
	return toModel(c), nil
}
