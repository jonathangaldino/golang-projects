package Env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func readEnvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

var MONGODB_URI = readEnvVariable("MONGODB_URI")
