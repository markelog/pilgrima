package database

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/markelog/pilgrima/database/models"
	"github.com/markelog/pilgrima/logger"
	"github.com/qor/validations"
)

// Up database
func Up() *gorm.DB {
	log := logger.Up()

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

	// Plugins
	validations.RegisterCallbacks(db)

	// Migrations
	db.AutoMigrate(
		&models.User{},
		&models.Project{},
		&models.Branch{},
		&models.Commit{},
		&models.Report{},
		&models.Token{},
	)

	return db
}
