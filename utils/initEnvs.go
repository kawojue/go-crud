package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv(path *string) {
	var err error

	if path == nil {
		err = godotenv.Load()
	} else {
		err = godotenv.Load(*path)
	}

	if err != nil {
		log.Fatal("Error loading .env file.")
	}
}

func GetEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
