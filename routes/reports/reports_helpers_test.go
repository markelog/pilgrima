package reports_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/markelog/pilgrima/database"
	"github.com/markelog/pilgrima/logger"
	"github.com/markelog/pilgrima/routes/reports"
	"github.com/markelog/pilgrima/test/env"
	"github.com/markelog/pilgrima/test/routes"
)

var (
	app *iris.Application
	db  *gorm.DB
)

func teardown() {
	db.Exec("TRUNCATE users CASCADE;")
	db.Exec("TRUNCATE projects CASCADE;")
	db.Exec("TRUNCATE tokens CASCADE;")
}

func TestMain(m *testing.M) {
	env.Up()

	app = routes.Up()
	db = database.Up()
	log := logger.Up()
	log.Out = ioutil.Discard

	reports.Up(app, db, log)

	app.Build()

	os.Exit(m.Run())
}