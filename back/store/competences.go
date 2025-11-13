package store

import (
	"back/constants"
	"back/model"
)

var competencesList = []model.Competence{
	{Name: "Go", Image: constants.GolangSVG, Rating: 5, Desc: constants.DescGO},
	{Name: "Java", Image: constants.JavaSVG, Rating: 3, Desc: constants.DescJava},
	{Name: "Spring", Image: constants.SpringSVG, Rating: 3, Desc: constants.DescSpring},
	{Name: "Python", Image: constants.PythonSVG, Rating: 4, Desc: constants.DescPython},
	{Name: "HTML", Image: constants.HtmlSVG, Rating: 5, Desc: constants.DescHTML},
	{Name: "CSS", Image: constants.CssSVG, Rating: 5, Desc: constants.DescCSS},
	{Name: "JavaScript", Image: constants.JsSVG, Rating: 4, Desc: constants.DescJS},
	{Name: "Svelte", Image: constants.SvelteSVG, Rating: 4, Desc: constants.DescSvelte},
	{Name: "TypeScript", Image: constants.TScriptSVG, Rating: 2, Desc: constants.DescTS},
	{Name: "Angular", Image: constants.AngularSVG, Rating: 3, Desc: constants.DescAngular},
	{Name: "SQL", Image: constants.SqlSVG, Rating: 4, Desc: constants.DescSQL},
	{Name: "Docker", Image: constants.DockerSVG, Rating: 4, Desc: constants.DescDocker},
	{Name: "Kubernetes", Image: constants.KubeSVG, Rating: 2, Desc: constants.DescKube},
	{Name: "Rebol", Image: constants.RebolSVG, Rating: 3, Desc: constants.DescRebol},
	{Name: "R", Image: constants.RStudioSVG, Rating: 3, Desc: constants.DescRStudio},
	{Name: "Ansible", Image: constants.AnsibleSVG, Rating: 2, Desc: constants.DescAnsible},
}

func (s *RAMStore) GetCompetences() (competences []model.Competence) {
	for i, competence := range competencesList {
		competence.ID = i + 1
		competences = append(competences, competence)
	}
	return competences
}
