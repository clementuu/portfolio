package cfg

import (
	"os"
)

const ErrInvConfig = "configuration invalide"

type Config struct {
	StaticDir string
	Host      string
	Port      string
}

var servConfig = Config{
	StaticDir: "./ihm",
}

func GetConfig() (Config, error) {
	// récupére le dossier statique
	dir := os.Getenv("DIR")
	if dir != "" {
		servConfig.StaticDir = dir
	}

	servConfig.Host = "127.0.0.1"
	servConfig.Port = ":8080"

	return servConfig, nil
}
