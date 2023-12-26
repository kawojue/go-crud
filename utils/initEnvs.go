package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv(paths ...*string) {
	var err error

	if paths == nil {
		err = godotenv.Load()
	} else {
		for _, path := range paths {
			err = godotenv.Load(*path)
			if err != nil {
				log.Fatal(fmt.Sprintf("Error loading %s file", *path), err)
			}
		}
	}

	if err != nil {
		log.Fatal("Error loading .env file.", err)
	}
}

func GetEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
