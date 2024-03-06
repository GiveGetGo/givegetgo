package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv(filepaths ...string) {
	fp := ".env"

	if len(filepaths) > 0 {
		fp = filepaths[0]
	}

	// Load environment variables
	err := godotenv.Load(fp)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}
