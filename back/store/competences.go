package store

import (
	"back/constants"
	"back/model"
)

var competencesList = []model.Competence{
	{Name: "Go", Image: constants.GolangSVG, Rating: 5, Desc: constants.DescGO, Type: model.Technique},
	{Name: "HTML", Image: constants.HtmlSVG, Rating: 5, Desc: constants.DescHTML, Type: model.Technique},
	{Name: "CSS", Image: constants.CssSVG, Rating: 5, Desc: constants.DescCSS, Type: model.Technique},
	{Name: "JavaScript", Image: constants.JsSVG, Rating: 4, Desc: constants.DescJS, Type: model.Technique},
	{Name: "Svelte", Image: constants.SvelteSVG, Rating: 4, Desc: constants.DescSvelte, Type: model.Technique},
	{Name: "TypeScript", Image: constants.TScriptSVG, Rating: 2, Desc: constants.DescTS, Type: model.Technique},
	{Name: "Angular", Image: constants.AngularSVG, Rating: 3, Desc: constants.DescAngular, Type: model.Technique},
	{Name: "SQL", Image: constants.SqlSVG, Rating: 4, Desc: constants.DescSQL, Type: model.Technique},
	{Name: "Java", Image: constants.JavaSVG, Rating: 2, Desc: constants.DescJava, Type: model.Technique},
	{Name: "Spring", Image: constants.SpringSVG, Rating: 3, Desc: constants.DescSpring, Type: model.Technique},
	{Name: "Python", Image: constants.PythonSVG, Rating: 3, Desc: constants.DescPython, Type: model.Technique},
	{Name: "Docker", Image: constants.DockerSVG, Rating: 4, Desc: constants.DescDocker, Type: model.Technique},
	{Name: "Kubernetes", Image: constants.KubeSVG, Rating: 2, Desc: constants.DescKube, Type: model.Technique},
	{Name: "Rebol", Image: constants.RebolSVG, Rating: 3, Desc: constants.DescRebol, Type: model.Technique},
	{Name: "R", Image: constants.RStudioSVG, Rating: 2, Desc: constants.DescRStudio, Type: model.Technique},
	{Name: "Ansible", Image: constants.AnsibleSVG, Rating: 2, Desc: constants.DescAnsible, Type: model.Technique},
	{Name: "Php", Image: constants.PhpSVG, Rating: 1, Desc: constants.DescPHP, Type: model.Technique},
	{Name: "Gestion de projet", Image: constants.ManageSVG, Rating: 2, Desc: "", Type: model.Humain},
	{Name: "Relation client", Image: constants.ClientSVG, Rating: 1, Desc: "", Type: model.Humain},
}

func (s *RAMStore) GetCompetences() []model.Competence {
	return competencesList
}
