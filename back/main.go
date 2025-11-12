package main

import (
	"back/cfg"
	"log"
	"net/http"
)

func main() {
	config, err := cfg.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	fs := http.FileServer(http.Dir(config.StaticDir))
	http.Handle("/", fs)

	log.Printf("Listening on %s...", config.Port)
	err = http.ListenAndServe(config.Port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
