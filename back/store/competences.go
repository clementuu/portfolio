package store

import (
	"back/constants"
	"back/model"
)

var (
	GoCompetence             = model.Competence{ID: 1, Name: "Go", Image: constants.GolangLogo, Rating: 4, Desc: constants.DescGO, Type: model.Technique}
	HTMLCompetence           = model.Competence{ID: 2, Name: "HTML", Image: constants.HtmlLogo, Rating: 4, Desc: constants.DescHTML, Type: model.Technique}
	CSSCompetence            = model.Competence{ID: 3, Name: "CSS", Image: constants.CssLogo, Rating: 4, Desc: constants.DescCSS, Type: model.Technique}
	JavaScriptCompetence     = model.Competence{ID: 4, Name: "JavaScript", Image: constants.JsLogo, Rating: 3, Desc: constants.DescJS, Type: model.Technique}
	SvelteCompetence         = model.Competence{ID: 5, Name: "Svelte", Image: constants.SvelteLogo, Rating: 4, Desc: constants.DescSvelte, Type: model.Technique}
	TypeScriptCompetence     = model.Competence{ID: 6, Name: "TypeScript", Image: constants.TScriptLogo, Rating: 2, Desc: constants.DescTS, Type: model.Technique}
	AngularCompetence        = model.Competence{ID: 7, Name: "Angular", Image: constants.AngularLogo, Rating: 2, Desc: constants.DescAngular, Type: model.Technique}
	SQLCompetence            = model.Competence{ID: 8, Name: "SQL", Image: constants.SqlLogo, Rating: 4, Desc: constants.DescSQL, Type: model.Technique}
	JavaCompetence           = model.Competence{ID: 9, Name: "Java", Image: constants.JavaLogo, Rating: 2, Desc: constants.DescJava, Type: model.Technique}
	SpringCompetence         = model.Competence{ID: 10, Name: "Spring", Image: constants.SpringLogo, Rating: 3, Desc: constants.DescSpring, Type: model.Technique}
	PythonCompetence         = model.Competence{ID: 11, Name: "Python", Image: constants.PythonLogo, Rating: 3, Desc: constants.DescPython, Type: model.Technique}
	DockerCompetence         = model.Competence{ID: 12, Name: "Docker", Image: constants.DockerLogo, Rating: 4, Desc: constants.DescDocker, Type: model.Technique}
	KubernetesCompetence     = model.Competence{ID: 13, Name: "Kubernetes", Image: constants.KubeLogo, Rating: 1, Desc: constants.DescKube, Type: model.Technique}
	RebolCompetence          = model.Competence{ID: 14, Name: "Rebol", Image: constants.RebolLogo, Rating: 2, Desc: constants.DescRebol, Type: model.Technique}
	RCompetence              = model.Competence{ID: 15, Name: "R", Image: constants.RLogo, Rating: 2, Desc: constants.DescRStudio, Type: model.Technique}
	AnsibleCompetence        = model.Competence{ID: 16, Name: "Ansible", Image: constants.AnsibleLogo, Rating: 1, Desc: constants.DescAnsible, Type: model.Technique}
	BashCompetence           = model.Competence{ID: 18, Name: "Bash", Image: constants.BashLogo, Rating: 2, Desc: "", Type: model.Technique}
	GestionProjetCompetence  = model.Competence{ID: 19, Name: "Gestion de projet", Image: constants.ManageLogo, Rating: 2, Desc: "", Type: model.Humain}
	RelationClientCompetence = model.Competence{ID: 20, Name: "Relation client", Image: constants.ClientLogo, Rating: 1, Desc: "", Type: model.Humain}
	FlexibiliteCompetence    = model.Competence{ID: 21, Name: "Flexibilité", Image: constants.FlexLogo, Rating: 4, Desc: "", Type: model.Humain}
	CreativiteCompetence     = model.Competence{ID: 22, Name: "Créativité", Image: constants.CreativeLogo, Rating: 3, Desc: "", Type: model.Humain}
	EspritCritiqueCompetence = model.Competence{ID: 23, Name: "Esprit critique", Image: constants.CriticalLogo, Rating: 4, Desc: "", Type: model.Humain}
)

var competencesList = []model.Competence{
	GoCompetence,
	HTMLCompetence,
	CSSCompetence,
	JavaScriptCompetence,
	SvelteCompetence,
	TypeScriptCompetence,
	AngularCompetence,
	SQLCompetence,
	JavaCompetence,
	SpringCompetence,
	PythonCompetence,
	DockerCompetence,
	KubernetesCompetence,
	RebolCompetence,
	RCompetence,
	AnsibleCompetence,
	BashCompetence,
	GestionProjetCompetence,
	RelationClientCompetence,
	FlexibiliteCompetence,
	CreativiteCompetence,
	EspritCritiqueCompetence,
}

func (s *RAMStore) GetCompetences() []model.Competence {
	return competencesList
}
