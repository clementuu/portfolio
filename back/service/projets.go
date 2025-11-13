package service

type Projet struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	Image       string       `json:"image"`
	Competences []Competence `json:"competences"`
	Desc        string       `json:"desc"`
}

func GetAllProjets() (projets []Projet) {
	ps := storage.GetProjets()
	for _, p := range ps {
		var projet Projet
		projet.ID = p.ID
		projet.Name = p.Name
		projet.Image = p.Image
		projet.Competences = getCompetences(p.Competences)
		projet.Desc = p.Desc
		projets = append(projets, projet)
	}
	return projets
}
