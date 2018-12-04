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
	var (
		log = logger.Up()
		dsn = os.Getenv("DATABASE_URL")

		err error
		db  *gorm.DB
	)

	if len(dsn) != 0 {
		db, err = models.ConnectDSN(dsn)
	} else {
		db, err = models.Connect(
			&models.ConnectArgs{
				User:     os.Getenv("DATABASE_USER"),
				Password: os.Getenv("DATABASE_PASSWORD"),
				Name:     os.Getenv("DATABASE_NAME"),
				Host:     os.Getenv("DATABASE_HOST"),
				Port:     os.Getenv("DATABASE_PORT"),
				SSL:      os.Getenv("DATABASE_SSL"),
			},
		)
	}
	if err != nil {
		log.Panic(err)
	}

	// Plugins
	validations.RegisterCallbacks(db)

	// Migrations
	err = db.AutoMigrate(
		&models.User{},
		&models.Project{},
		&models.Branch{},
		&models.Commit{},
		&models.Report{},
		&models.Token{},
	).Error

	if err != nil {
		log.Panic(err)
	}

	return db
}
