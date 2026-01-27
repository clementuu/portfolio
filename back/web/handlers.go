package web

import (
	"back/service/competence"
	"back/service/cv"
	"back/service/projet"
	"net/http"
	"strconv"
)

// GetAllCompetences renvoie la liste des compétences
func GetAllCompetences(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	err := encodeJSON(w, competence.GetAll())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetCompetenceByID renvoie une compétence d'id donné
func GetCompetenceByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var compentence competence.Model
	compentence, err = competence.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

	err = encodeJSON(w, compentence)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetAllProjets renvoie la liste des projets
func GetAllProjets(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	err := encodeJSON(w, projet.GetAll())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetProjetsNames renvoie les noms et ID des différents projets
func GetProjetsNames(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	err := encodeJSON(w, projet.GetNames())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetProjetByID renvoie un projet d'id donné
func GetProjetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var proj projet.Model
	proj, err = projet.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

	err = encodeJSON(w, proj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetFormations renvoie la liste des formations
func GetFormations(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	err := encodeJSON(w, cv.GetFormations())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetExperiences renvoie la liste des expériences professionnelles
func GetExperiences(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	err := encodeJSON(w, cv.GetExperiences())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
