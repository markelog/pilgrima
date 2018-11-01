package database

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/markelog/pilgrima/database/models"
)

// Up database
func Up() *gorm.DB {
	db, err := models.Connect(
		&models.ConnectArgs{
			User:     os.Getenv("DATABASE_USER"),
			Password: os.Getenv("DATABASE_PASSWORD"),
			Name:     os.Getenv("DATABASE_NAME"),
			Host:     os.Getenv("DATABASE_HOST"),
			Port:     os.Getenv("DATABASE_PORT"),
			SSL:      os.Getenv("DATABASE_SSL"),
		},
	)
	if err != nil {
		log.Panic(err)
	}

	return db
}
