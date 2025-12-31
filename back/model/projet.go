package model

type Projet struct {
	ID          int
	Name        string
	Image       string
	Competences []Competence
	Resume      string
	Template    string
}

type MiniProjet struct {
	ID   int
	Name string
}
