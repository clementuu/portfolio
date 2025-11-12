package web

import (
	"back/service"
	"net/http"
)

// GetCompetences envoie la liste des compétences
func GetCompetences(w http.ResponseWriter, r *http.Request) {
	err := encodeJSON(w, service.GetCompetences())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
