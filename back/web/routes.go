package web

import "net/http"

func MakeRoutes() {
	routes := []struct {
		pattern string
		handler http.HandlerFunc
	}{
		{"GET /competences", GetAllCompetences},
		{"GET /projets", GetAllProjets},
	}

	for _, r := range routes {
		http.HandleFunc(r.pattern, r.handler)
	}
}
