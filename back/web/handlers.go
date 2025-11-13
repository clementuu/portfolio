package web

import (
	"back/service"
	"net/http"
)

// GetAllCompetences envoie la liste des compétences
func GetAllCompetences(w http.ResponseWriter, r *http.Request) {
	err := encodeJSON(w, service.GetAllCompetences())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetAllProjets renvoie la liste des projets
func GetAllProjets(w http.ResponseWriter, r *http.Request) {
	err := encodeJSON(w, service.GetAllProjets())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
