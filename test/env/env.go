package env

import (
	"log"

	"github.com/joho/godotenv"
)

// Up test environments
func Up() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Panic(err)
	}
}
