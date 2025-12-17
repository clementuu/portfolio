package web

import "net/http"

func MakeRoutes() {
	routes := []struct {
		pattern string
		handler http.HandlerFunc
	}{
		{"GET /competences", GetAllCompetences},
		{"GET /projets", GetAllProjets},
		{"GET /projets/names", GetProjetsNames},
		{"GET /projet/{id}", GetProjetByID},
		{"GET /formations", GetFormations},
	}

	for _, r := range routes {
		http.HandleFunc(r.pattern, r.handler)
	}
}
