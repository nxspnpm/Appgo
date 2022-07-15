package config

import (
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	DatabaseDSN string
}

var (
	AppConfig *config
)

func LoadConfig() {
	godotenv.Load()

	AppConfig = &config{
		DatabaseDSN: os.Getenv("DATABASE_DSN"),
	}
}
