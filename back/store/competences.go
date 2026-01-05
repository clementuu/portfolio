package store

import (
	"back/constants"
	"back/model"
	tmpl "back/tmpl/competences"
	"fmt"
)

var ErrNoCompetence = "aucune compétence trouvée pour l'id %d"

var (
	GoCompetence             = model.Competence{ID: 1, Name: "Go", Image: constants.GolangLogo, Rating: 4, Template: tmpl.GoComp, Type: model.Dev}
	WebCompetence            = model.Competence{ID: 2, Name: "HTML/CSS", Image: constants.WebLogo, Rating: 4, Template: tmpl.WebComp, Type: model.Dev}
	JavaScriptCompetence     = model.Competence{ID: 3, Name: "JavaScript", Image: constants.JsLogo, Rating: 3, Template: tmpl.JavaScriptComp, Type: model.Dev}
	SvelteCompetence         = model.Competence{ID: 4, Name: "Svelte", Image: constants.SvelteLogo, Rating: 4, Template: tmpl.SvelteComp, Type: model.Dev}
	AngularCompetence        = model.Competence{ID: 5, Name: "Angular", Image: constants.AngularLogo, Rating: 2, Template: tmpl.AngularComp, Type: model.Dev}
	SQLCompetence            = model.Competence{ID: 6, Name: "SQL", Image: constants.SqlLogo, Rating: 4, Template: tmpl.SQLComp, Type: model.Dev}
	JavaCompetence           = model.Competence{ID: 7, Name: "Java", Image: constants.JavaLogo, Rating: 2, Template: tmpl.JavaComp, Type: model.Dev}
	SpringCompetence         = model.Competence{ID: 8, Name: "Spring", Image: constants.SpringLogo, Rating: 3, Template: tmpl.SpringComp, Type: model.Dev}
	PythonCompetence         = model.Competence{ID: 9, Name: "Python", Image: constants.PythonLogo, Rating: 3, Template: tmpl.PythonComp, Type: model.Dev}
	RebolCompetence          = model.Competence{ID: 10, Name: "Rebol", Image: constants.RebolLogo, Rating: 2, Template: tmpl.RebolComp, Type: model.Dev}
	RCompetence              = model.Competence{ID: 11, Name: "R", Image: constants.RLogo, Rating: 2, Template: tmpl.RComp, Type: model.Dev}
	DockerCompetence         = model.Competence{ID: 12, Name: "Docker", Image: constants.DockerLogo, Rating: 4, Template: tmpl.DockerComp, Type: model.DevOps}
	KubernetesCompetence     = model.Competence{ID: 13, Name: "Kubernetes", Image: constants.KubeLogo, Rating: 1, Template: tmpl.KubernetesComp, Type: model.DevOps}
	BashCompetence           = model.Competence{ID: 14, Name: "Bash", Image: constants.BashLogo, Rating: 2, Template: tmpl.BashComp, Type: model.DevOps}
	GestionProjetCompetence  = model.Competence{ID: 15, Name: "Gestion de projet", Image: constants.ManageLogo, Rating: 2, Template: tmpl.GestionProjetComp, Type: model.Human}
	FlexibiliteCompetence    = model.Competence{ID: 16, Name: "Flexibilité", Image: constants.FlexLogo, Rating: 4, Template: tmpl.FlexibiliteComp, Type: model.Human}
	CreativiteCompetence     = model.Competence{ID: 17, Name: "Créativité", Image: constants.CreativeLogo, Rating: 3, Template: tmpl.CreativiteComp, Type: model.Human}
	EspritCritiqueCompetence = model.Competence{ID: 18, Name: "Esprit critique", Image: constants.CriticalLogo, Rating: 4, Template: tmpl.PenseeCritiqueComp, Type: model.Human}
)

var competencesList = []model.Competence{
	GoCompetence,
	WebCompetence,
	JavaScriptCompetence,
	SvelteCompetence,
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

func (s *RAMStore) GetCompetence(id int) (model.Competence, error) {
	for _, c := range competencesList {
		if c.ID == id {
			return c, nil
		}
	}
	return model.Competence{}, fmt.Errorf(ErrNoCompetence, id)
}
