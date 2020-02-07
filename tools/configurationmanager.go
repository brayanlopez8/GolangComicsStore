package tools

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

//ReadConfig read key config by name
func ReadConfig(key string) string {

	// load .env file
	err := godotenv.Load(".envdev")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
