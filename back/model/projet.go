package model

type Projet struct {
	ID          int
	Name        string
	Image       string
	Competences []Competence
	Desc        string
}
