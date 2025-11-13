package model

type CompetenceType int

const (
	Technique CompetenceType = 1
	Humain    CompetenceType = 2
)

type Competence struct {
	ID     int
	Name   string
	Image  string
	Rating int
	Desc   string
	Exp    string
	Type   CompetenceType
}

func (ct CompetenceType) String() string {
	switch ct {
	case Technique:
		return "Technique"
	case Humain:
		return "Humain"
	default:
		return "Type inconnu"
	}
}
