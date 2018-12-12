package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Up test environment
func Up() {
	err := godotenv.Load("../../.env")

	// Always ignore fixtures for tests - its tests responsobility to fill fixtures
	os.Setenv("DATABASE_FIXTURES_PATH", "")

	if err != nil {
		log.Panic(err)
	}
}
