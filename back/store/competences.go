package store

import (
	"back/constants"
	"back/model"
	tmpl "back/tmpl/competences"
	"fmt"
)

var ErrNoCompetence = "aucune compétence trouvée pour l'id %d"

var (
	goCompetence             = model.Competence{ID: 1, Name: "Go", Image: constants.GolangLogo, Rating: 4, Template: tmpl.GoCompetence, Type: model.Dev}
	webCompetence            = model.Competence{ID: 2, Name: "HTML/CSS", Image: constants.WebLogo, Rating: 4, Template: tmpl.WebCompetence, Type: model.Dev}
	javaScriptCompetence     = model.Competence{ID: 3, Name: "JavaScript", Image: constants.JsLogo, Rating: 3, Template: tmpl.JavaScriptCompetence, Type: model.Dev}
	svelteCompetence         = model.Competence{ID: 4, Name: "Svelte", Image: constants.SvelteLogo, Rating: 4, Template: tmpl.SvelteCompetence, Type: model.Dev}
	angularCompetence        = model.Competence{ID: 5, Name: "Angular", Image: constants.AngularLogo, Rating: 2, Template: tmpl.AngularCompetence, Type: model.Dev}
	sqlCompetence            = model.Competence{ID: 6, Name: "SQL", Image: constants.SqlLogo, Rating: 4, Template: tmpl.SQLCompetence, Type: model.Dev}
	springCompetence         = model.Competence{ID: 7, Name: "Spring", Image: constants.SpringLogo, Rating: 3, Template: tmpl.SpringCompetence, Type: model.Dev}
	pythonCompetence         = model.Competence{ID: 8, Name: "Python", Image: constants.PythonLogo, Rating: 3, Template: tmpl.PythonCompetence, Type: model.Dev}
	rebolCompetence          = model.Competence{ID: 9, Name: "Rebol", Image: constants.RebolLogo, Rating: 2, Template: tmpl.RebolCompetence, Type: model.Dev}
	rCompetence              = model.Competence{ID: 10, Name: "R", Image: constants.RLogo, Rating: 2, Template: tmpl.RCompetence, Type: model.Dev}
	dockerCompetence         = model.Competence{ID: 11, Name: "Docker", Image: constants.DockerLogo, Rating: 4, Template: tmpl.DockerCompetence, Type: model.DevOps}
	bashCompetence           = model.Competence{ID: 12, Name: "Bash", Image: constants.BashLogo, Rating: 2, Template: tmpl.BashCompetence, Type: model.DevOps}
	gestionProjetCompetence  = model.Competence{ID: 13, Name: "Gestion de projet", Image: constants.ManageLogo, Rating: 2, Template: tmpl.GestionProjetCompetence, Type: model.Human}
	flexibiliteCompetence    = model.Competence{ID: 14, Name: "Flexibilité", Image: constants.FlexLogo, Rating: 4, Template: tmpl.FlexibiliteCompetence, Type: model.Human}
	creativiteCompetence     = model.Competence{ID: 15, Name: "Créativité", Image: constants.CreativeLogo, Rating: 3, Template: tmpl.CreativiteCompetence, Type: model.Human}
	espritCritiqueCompetence = model.Competence{ID: 16, Name: "Esprit critique", Image: constants.CriticalLogo, Rating: 4, Template: tmpl.EspritCritiqueCompetence, Type: model.Human}
)

var competencesList = []model.Competence{
	goCompetence,
	webCompetence,
	javaScriptCompetence,
	svelteCompetence,
	angularCompetence,
	sqlCompetence,
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
