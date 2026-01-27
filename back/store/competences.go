package store

import (
	"back/constants"
	"back/model"
	tmpl "back/tmpl/competences"
	"fmt"
)

var ErrNoCompetence = "aucune compétence trouvée pour l'id %d"

var (
	goCompetence             = model.Competence{ID: 1, Name: "Go", Image: constants.GolangLogo, Rating: 4, Template: tmpl.GoComp, Type: model.Dev}
	webCompetence            = model.Competence{ID: 2, Name: "HTML/CSS", Image: constants.WebLogo, Rating: 4, Template: tmpl.WebComp, Type: model.Dev}
	javaScriptCompetence     = model.Competence{ID: 3, Name: "JavaScript", Image: constants.JsLogo, Rating: 3, Template: tmpl.JavaScriptComp, Type: model.Dev}
	svelteCompetence         = model.Competence{ID: 4, Name: "Svelte", Image: constants.SvelteLogo, Rating: 4, Template: tmpl.SvelteComp, Type: model.Dev}
	angularCompetence        = model.Competence{ID: 5, Name: "Angular", Image: constants.AngularLogo, Rating: 2, Template: tmpl.AngularComp, Type: model.Dev}
	sqlCompetence            = model.Competence{ID: 6, Name: "SQL", Image: constants.SqlLogo, Rating: 4, Template: tmpl.SQLComp, Type: model.Dev}
	javaCompetence           = model.Competence{ID: 7, Name: "Java", Image: constants.JavaLogo, Rating: 2, Template: tmpl.JavaComp, Type: model.Dev}
	springCompetence         = model.Competence{ID: 8, Name: "Spring", Image: constants.SpringLogo, Rating: 3, Template: tmpl.SpringComp, Type: model.Dev}
	pythonCompetence         = model.Competence{ID: 9, Name: "Python", Image: constants.PythonLogo, Rating: 3, Template: tmpl.PythonComp, Type: model.Dev}
	rebolCompetence          = model.Competence{ID: 10, Name: "Rebol", Image: constants.RebolLogo, Rating: 2, Template: tmpl.RebolComp, Type: model.Dev}
	rCompetence              = model.Competence{ID: 11, Name: "R", Image: constants.RLogo, Rating: 2, Template: tmpl.RComp, Type: model.Dev}
	dockerCompetence         = model.Competence{ID: 12, Name: "Docker", Image: constants.DockerLogo, Rating: 4, Template: tmpl.DockerComp, Type: model.DevOps}
	bashCompetence           = model.Competence{ID: 13, Name: "Bash", Image: constants.BashLogo, Rating: 2, Template: tmpl.BashComp, Type: model.DevOps}
	gestionProjetCompetence  = model.Competence{ID: 14, Name: "Gestion de projet", Image: constants.ManageLogo, Rating: 2, Template: tmpl.GestionProjetComp, Type: model.Human}
	flexibiliteCompetence    = model.Competence{ID: 15, Name: "Flexibilité", Image: constants.FlexLogo, Rating: 4, Template: tmpl.FlexibiliteComp, Type: model.Human}
	creativiteCompetence     = model.Competence{ID: 16, Name: "Créativité", Image: constants.CreativeLogo, Rating: 3, Template: tmpl.CreativiteComp, Type: model.Human}
	espritCritiqueCompetence = model.Competence{ID: 17, Name: "Esprit critique", Image: constants.CriticalLogo, Rating: 4, Template: tmpl.PenseeCritiqueComp, Type: model.Human}
)

var competencesList = []model.Competence{
	goCompetence,
	webCompetence,
	javaScriptCompetence,
	svelteCompetence,
	angularCompetence,
	sqlCompetence,
	javaCompetence,
	springCompetence,
	pythonCompetence,
	rebolCompetence,
	rCompetence,
	dockerCompetence,
	bashCompetence,
	gestionProjetCompetence,
	flexibiliteCompetence,
	creativiteCompetence,
	espritCritiqueCompetence,
}

func (r *RAMStore) GetCompetences() []model.Competence {
	return r.Competences
}

func (r *RAMStore) GetCompetence(id int) (model.Competence, error) {
	for _, c := range r.Competences {
		if c.ID == id {
			return c, nil
		}
	}
	return model.Competence{}, fmt.Errorf(ErrNoCompetence, id)
}
