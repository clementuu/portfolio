package main

import (
	"back/cfg"
	"back/service"
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

	ramStore := store.NewRAMStore()

	service.Init(ramStore)

	fs := http.FileServer(http.Dir(config.StaticDir))
	http.Handle("/", fs)

	web.MakeRoutes()

	log.Printf("Listening on %s...", config.Port)
	err = http.ListenAndServe(config.Port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
