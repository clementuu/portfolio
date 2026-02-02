package web

import (
	"back/log"
	"back/service/competence"
	"back/service/cv"
	"back/service/projet"
	"net/http"
	"strconv"
)

// Handler struct holds application-wide dependencies like the logger.
type Handler struct {
	Logger log.Logger
}

// GetAllCompetences renvoie la liste des compétences
func (h *Handler) GetAllCompetences(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	err := encodeJSON(w, competence.GetAll())
	if err != nil {
		h.Logger.Error("Error encoding competences: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetCompetenceByID renvoie une compétence d'id donné
func (h *Handler) GetCompetenceByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.Logger.Warn("Invalid ID format: %s", idStr)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var compentence competence.Model
	compentence, err = competence.Get(id)
	if err != nil {
		h.Logger.Error("Error getting competence by ID %d: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

	err = encodeJSON(w, compentence)
	if err != nil {
		h.Logger.Error("Error encoding competence: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetAllProjets renvoie la liste des projets
func (h *Handler) GetAllProjets(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	err := encodeJSON(w, projet.GetAll())
	if err != nil {
		h.Logger.Error("Error encoding projets: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetProjetsNames renvoie les noms et ID des différents projets
func (h *Handler) GetProjetsNames(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	err := encodeJSON(w, projet.GetNames())
	if err != nil {
		h.Logger.Error("Error encoding project names: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetProjetByID renvoie un projet d'id donné
func (h *Handler) GetProjetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.Logger.Warn("Invalid ID format: %s", idStr)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var proj projet.Model
	proj, err = projet.Get(id)
	if err != nil {
		h.Logger.Error("Error getting projet by ID %d: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

	err = encodeJSON(w, proj)
	if err != nil {
		h.Logger.Error("Error encoding projet: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetFormations renvoie la liste des formations
func (h *Handler) GetFormations(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	err := encodeJSON(w, cv.GetFormations())
	if err != nil {
		h.Logger.Error("Error encoding formations: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetExperiences renvoie la liste des expériences professionnelles
func (h *Handler) GetExperiences(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	err := encodeJSON(w, cv.GetExperiences())
	if err != nil {
		h.Logger.Error("Error encoding experiences: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// LoggingMiddleware logs incoming HTTP requests.
func LoggingMiddleware(logger log.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info("%s %s %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
