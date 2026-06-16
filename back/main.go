package main

import (
	"back/cfg"
	"back/service/blog"
	"back/service/competence"
	"back/service/cv"
	"back/service/projet"
	"back/store"
	"back/web"
	"log"
	"net/http"
)

func main() {
	config, err := cfg.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Initialise le RAMStore et configure les services avec celui-ci
	ramStore := store.NewRAMStore()

	competence.Setup(ramStore)
	projet.Setup(ramStore)
	cv.Setup(ramStore)
	blog.Setup(ramStore)

	fs := http.FileServer(http.Dir(config.StaticDir))
	http.Handle("/", fs)

	web.MakeRoutes()

	log.Printf("Listening on %s...", config.Port)
	err = http.ListenAndServe(config.Port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
