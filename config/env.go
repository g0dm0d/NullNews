package config

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Env struct {
	ApiAddr  string `env:"API_Addr"`
	ApiPort  string `env:"API_PORT"`
	DBDSN    string `env:"DB_DSN"`
	DBDriver string `env:"DB_Dirver"`
	Secret   string `env:"SECRET"`
}

func envParse() *Env {
	var e Env
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}
	err = env.Parse(&e)
	if err != nil {
		log.Fatal("Error parse .env file")
	}
	return &e
}
