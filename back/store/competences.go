package store

import "back/model"

var competencesList = []model.Competence{
	{Name: "Go", Image: "../assets/icons8-golang.svg", Rating: 5},
	{Name: "Java", Image: "../assets/icons8-logo-java-coffee-cup.svg", Rating: 3},
	{Name: "Spring", Image: "../assets/icons8-logo-de-printemps.svg", Rating: 3},
	{Name: "HTML", Image: "../assets/icons8-html.svg", Rating: 5},
	{Name: "CSS", Image: "../assets/icons8-css.svg", Rating: 5},
	{Name: "JavaScript", Image: "../assets/icons8-js.svg", Rating: 4},
	{Name: "Svelte", Image: "../assets/Svelte_Logo.svg", Rating: 4},
	{Name: "SQL", Image: "../assets/icons8-mysql.svg", Rating: 4},
	{Name: "Docker", Image: "../assets/icons8-docker.svg", Rating: 4},
	{Name: "Kubernetes", Image: "../assets/icons8-kubernetes.svg", Rating: 2},
	{Name: "Rebol", Image: "../assets/Rebol_logo.png", Rating: 3},
	{Name: "R", Image: "../assets/RStudio_logo_flat.svg", Rating: 3},
}

func (s *RAMStore) GetCompetences() []model.Competence {
	return competencesList
}
