package env

import (
	"github.com/markelog/pilgrima/logger"

	"github.com/joho/godotenv"
)

// Up environment
func Up() {
	log := logger.Up()

	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
	}
}
