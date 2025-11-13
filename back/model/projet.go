package model

type Projet struct {
	ID          int
	Name        string
	Image       string
	Competences []Competence
	Desc        string
}

type MiniProjet struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
