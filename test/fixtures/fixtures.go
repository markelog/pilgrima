package fixtures

import (
	"github.com/jinzhu/gorm"
	"github.com/markelog/pilgrima/logger"
	testfixtures "gopkg.in/testfixtures.v2"
)

var (
	db       *gorm.DB
	fixtures *testfixtures.Context
)

// Up fixtures
func Up(path string, db *gorm.DB) *testfixtures.Context {
	log := logger.Up()

	fixtures, err := testfixtures.NewFolder(
		db.DB(),
		&testfixtures.PostgreSQL{},
		path,
	)
	if err != nil {
		log.Fatal(err)
	}

	return fixtures
}
