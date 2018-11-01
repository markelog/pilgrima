package fixtures

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/markelog/pilgrima/database/models"
	testfixtures "gopkg.in/testfixtures.v2"
)

func Up() {
	var (
		db       *gorm.DB
		fixtures *testfixtures.Context
	)

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

	fixtures, err = testfixtures.NewFolder(db.DB(), &testfixtures.PostgreSQL{}, "database/fixtures/development")
	if err != nil {
		log.Fatal(err)
	}

	if err := fixtures.Load(); err != nil {
		log.Fatal(err)
	}
}
