package fixtures

import (
	"log"

	"github.com/jinzhu/gorm"
	testfixtures "gopkg.in/testfixtures.v2"
)

var (
	db       *gorm.DB
	fixtures *testfixtures.Context
)

// Up fixtures
func Up(path string, db *gorm.DB) *testfixtures.Context {
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
