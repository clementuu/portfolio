package web

import (
	"back/log"
	"net/http"
)

// MakeRoutes initializes all the routes for the application.
func MakeRoutes() {
	logger := log.NewLogger()
	handler := &Handler{Logger: logger}

	routes := []struct {
		pattern string
		handler http.HandlerFunc
	}{
		{"GET /competences", handler.GetAllCompetences},
		{"GET /competence/{id}", handler.GetCompetenceByID},
		{"GET /projets", handler.GetAllProjets},
		{"GET /projets/names", handler.GetProjetsNames},
		{"GET /projet/{id}", handler.GetProjetByID},
		{"GET /formations", handler.GetFormations},
		{"GET /experiences", handler.GetExperiences},
	}

	for _, r := range routes {
		http.Handle(r.pattern, LoggingMiddleware(logger, r.handler))
	}
}
