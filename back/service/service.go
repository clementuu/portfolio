package service

type Competence struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Image  string `json:"image"`
	Rating int    `json:"rating"`
	Desc   string `json:"desc"`
	Exp    string `json:"exp"`
	Type   string `json:"type"`
}

func GetCompetences() (competences []Competence) {
	cs := storage.GetCompetences()
	for i, c := range cs {
		var competence Competence
		competence.ID = i + 1
		competence.Name = c.Name
		competence.Image = c.Image
		competence.Rating = c.Rating
		competence.Desc = c.Desc
		competence.Exp = c.Exp
		competence.Type = c.Type.String()
		competences = append(competences, competence)
	}
	return competences
}
