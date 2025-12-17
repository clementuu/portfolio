package web

import (
	"back/service"
	"net/http"
	"strconv"
)

// GetAllCompetences renvoie la liste des compétences
func GetAllCompetences(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	err := encodeJSON(w, service.GetAllCompetences())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetAllProjets renvoie la liste des projets
func GetAllProjets(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	err := encodeJSON(w, service.GetAllProjets())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetProjetsNames renvoie les noms et ID des différents projets
func GetProjetsNames(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	err := encodeJSON(w, service.GetProjetsNames())
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
	var projet service.Projet
	projet, err = service.GetProjet(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

	err = encodeJSON(w, projet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetFormations renvoie la liste des formations
func GetFormations(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	err := encodeJSON(w, service.GetFormations())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
