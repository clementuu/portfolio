package model

type CompetenceType int

const (
	Dev    CompetenceType = 1
	DevOps CompetenceType = 2
	Human  CompetenceType = 3
)

type Competence struct {
	ID     int
	Name   string
	Image  string
	Rating int
	Template   string
	Exp    string
	Type   CompetenceType
}

func (ct CompetenceType) String() string {
	switch ct {
	case Dev:
		return "Dev"
	case DevOps:
		return "DevOps"
	case Human:
		return "Humain"
	default:
		return "Type inconnu"
	}
}
