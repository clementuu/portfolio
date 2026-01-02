package store

import (
	"back/constants"
	"back/model"
)

var (
	GoCompetence             = model.Competence{ID: 1, Name: "Go", Image: constants.GolangLogo, Rating: 4, Desc: constants.DescGO, Type: model.Dev}
	HTMLCompetence           = model.Competence{ID: 2, Name: "HTML", Image: constants.HtmlLogo, Rating: 4, Desc: constants.DescHTML, Type: model.Dev}
	CSSCompetence            = model.Competence{ID: 3, Name: "CSS", Image: constants.CssLogo, Rating: 4, Desc: constants.DescCSS, Type: model.Dev}
	JavaScriptCompetence     = model.Competence{ID: 4, Name: "JavaScript", Image: constants.JsLogo, Rating: 3, Desc: constants.DescJS, Type: model.Dev}
	SvelteCompetence         = model.Competence{ID: 5, Name: "Svelte", Image: constants.SvelteLogo, Rating: 4, Desc: constants.DescSvelte, Type: model.Dev}
	TypeScriptCompetence     = model.Competence{ID: 6, Name: "TypeScript", Image: constants.TScriptLogo, Rating: 2, Desc: constants.DescTS, Type: model.Dev}
	AngularCompetence        = model.Competence{ID: 7, Name: "Angular", Image: constants.AngularLogo, Rating: 2, Desc: constants.DescAngular, Type: model.Dev}
	SQLCompetence            = model.Competence{ID: 8, Name: "SQL", Image: constants.SqlLogo, Rating: 4, Desc: constants.DescSQL, Type: model.Dev}
	JavaCompetence           = model.Competence{ID: 9, Name: "Java", Image: constants.JavaLogo, Rating: 2, Desc: constants.DescJava, Type: model.Dev}
	SpringCompetence         = model.Competence{ID: 10, Name: "Spring", Image: constants.SpringLogo, Rating: 3, Desc: constants.DescSpring, Type: model.Dev}
	PythonCompetence         = model.Competence{ID: 11, Name: "Python", Image: constants.PythonLogo, Rating: 3, Desc: constants.DescPython, Type: model.Dev}
	RebolCompetence          = model.Competence{ID: 14, Name: "Rebol", Image: constants.RebolLogo, Rating: 2, Desc: constants.DescRebol, Type: model.Dev}
	RCompetence              = model.Competence{ID: 15, Name: "R", Image: constants.RLogo, Rating: 2, Desc: constants.DescRStudio, Type: model.Dev}
	DockerCompetence         = model.Competence{ID: 12, Name: "Docker", Image: constants.DockerLogo, Rating: 4, Desc: constants.DescDocker, Type: model.DevOps}
	KubernetesCompetence     = model.Competence{ID: 13, Name: "Kubernetes", Image: constants.KubeLogo, Rating: 1, Desc: constants.DescKube, Type: model.DevOps}
	BashCompetence           = model.Competence{ID: 18, Name: "Bash", Image: constants.BashLogo, Rating: 2, Desc: "", Type: model.DevOps}
	GestionProjetCompetence  = model.Competence{ID: 19, Name: "Gestion de projet", Image: constants.ManageLogo, Rating: 2, Desc: "", Type: model.Human}
	FlexibiliteCompetence    = model.Competence{ID: 21, Name: "Flexibilité", Image: constants.FlexLogo, Rating: 4, Desc: "", Type: model.Human}
	CreativiteCompetence     = model.Competence{ID: 22, Name: "Créativité", Image: constants.CreativeLogo, Rating: 3, Desc: "", Type: model.Human}
	EspritCritiqueCompetence = model.Competence{ID: 23, Name: "Esprit critique", Image: constants.CriticalLogo, Rating: 4, Desc: "", Type: model.Human}
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
	RebolCompetence,
	RCompetence,
	DockerCompetence,
	KubernetesCompetence,
	BashCompetence,
	GestionProjetCompetence,
	FlexibiliteCompetence,
	CreativiteCompetence,
	EspritCritiqueCompetence,
}

func (s *RAMStore) GetCompetences() []model.Competence {
	return competencesList
}
