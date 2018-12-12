package env

import (
	"os"
	"path"

	"github.com/jinzhu/gorm"
	"github.com/markelog/pilgrima/logger"
	testfixtures "gopkg.in/testfixtures.v2"

	"github.com/joho/godotenv"
)

// Up environment
func Up() {
	log := logger.Up()

	err := godotenv.Load()
	if err != nil {
		log.Info("Haven't load the .env file")
	}
}

// Fixtures load them
func Fixtures(relative string, db *gorm.DB) (*testfixtures.Context, error) {
	dir, _ := os.Getwd()
	absolute := path.Join(dir, relative)

	fixtures, err := testfixtures.NewFolder(
		db.DB(),
		&testfixtures.PostgreSQL{},
		absolute,
	)
	if err != nil {
		return nil, err
	}

	err = fixtures.Load()

	return fixtures, err
}
