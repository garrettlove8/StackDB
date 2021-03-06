package utils

import (
	"log"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

// GetEnv gets the correct environment variables.
// It is executed at the very beginning of running the database.
func GetEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
}

func GetUuid() string {
	return uuid.New().String()
}
