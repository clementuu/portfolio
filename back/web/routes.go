package web

import "net/http"

func MakeRoutes() {
	routes := []struct {
		pattern string
		handler http.HandlerFunc
	}{
		{"GET /competences", GetCompetences},
	}

	for _, r := range routes {
		http.HandleFunc(r.pattern, r.handler)
	}
}
