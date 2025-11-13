package store

import (
	"back/constants"
	"back/model"
)

var GoCompetence = model.Competence{ID: 1, Name: "Go", Image: constants.GolangLogo, Rating: 5, Desc: constants.DescGO, Type: model.Technique}
var HTMLCompetence = model.Competence{ID: 2, Name: "HTML", Image: constants.HtmlLogo, Rating: 5, Desc: constants.DescHTML, Type: model.Technique}
var CSSCompetence = model.Competence{ID: 3, Name: "CSS", Image: constants.CssLogo, Rating: 5, Desc: constants.DescCSS, Type: model.Technique}
var JavaScriptCompetence = model.Competence{ID: 4, Name: "JavaScript", Image: constants.JsLogo, Rating: 4, Desc: constants.DescJS, Type: model.Technique}
var SvelteCompetence = model.Competence{ID: 5, Name: "Svelte", Image: constants.SvelteLogo, Rating: 4, Desc: constants.DescSvelte, Type: model.Technique}
var TypeScriptCompetence = model.Competence{ID: 6, Name: "TypeScript", Image: constants.TScriptLogo, Rating: 2, Desc: constants.DescTS, Type: model.Technique}
var AngularCompetence = model.Competence{ID: 7, Name: "Angular", Image: constants.AngularLogo, Rating: 3, Desc: constants.DescAngular, Type: model.Technique}
var SQLCompetence = model.Competence{ID: 8, Name: "SQL", Image: constants.SqlLogo, Rating: 4, Desc: constants.DescSQL, Type: model.Technique}
var JavaCompetence = model.Competence{ID: 9, Name: "Java", Image: constants.JavaLogo, Rating: 2, Desc: constants.DescJava, Type: model.Technique}
var SpringCompetence = model.Competence{ID: 10, Name: "Spring", Image: constants.SpringLogo, Rating: 3, Desc: constants.DescSpring, Type: model.Technique}
var PythonCompetence = model.Competence{ID: 11, Name: "Python", Image: constants.PythonLogo, Rating: 3, Desc: constants.DescPython, Type: model.Technique}
var DockerCompetence = model.Competence{ID: 12, Name: "Docker", Image: constants.DockerLogo, Rating: 4, Desc: constants.DescDocker, Type: model.Technique}
var KubernetesCompetence = model.Competence{ID: 13, Name: "Kubernetes", Image: constants.KubeLogo, Rating: 2, Desc: constants.DescKube, Type: model.Technique}
var RebolCompetence = model.Competence{ID: 14, Name: "Rebol", Image: constants.RebolLogo, Rating: 3, Desc: constants.DescRebol, Type: model.Technique}
var RCompetence = model.Competence{ID: 15, Name: "R", Image: constants.RLogo, Rating: 2, Desc: constants.DescRStudio, Type: model.Technique}
var AnsibleCompetence = model.Competence{ID: 16, Name: "Ansible", Image: constants.AnsibleLogo, Rating: 2, Desc: constants.DescAnsible, Type: model.Technique}
var PhpCompetence = model.Competence{ID: 17, Name: "Php", Image: constants.PhpLogo, Rating: 1, Desc: constants.DescPHP, Type: model.Technique}
var GestionProjetCompetence = model.Competence{ID: 18, Name: "Gestion de projet", Image: constants.ManageLogo, Rating: 2, Desc: "", Type: model.Humain}
var RelationClientCompetence = model.Competence{ID: 19, Name: "Relation client", Image: constants.ClientLogo, Rating: 1, Desc: "", Type: model.Humain}

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
	PhpCompetence,
	GestionProjetCompetence,
	RelationClientCompetence,
}

func (s *RAMStore) GetCompetences() []model.Competence {
	return competencesList
}
