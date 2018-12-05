package token_test

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/markelog/pilgrima/database"
	"github.com/markelog/pilgrima/logger"
	"github.com/markelog/pilgrima/routes/token"
	"github.com/markelog/pilgrima/test/env"
	"github.com/markelog/pilgrima/test/fixtures"
	"github.com/markelog/pilgrima/test/request"
	"github.com/markelog/pilgrima/test/routes"
	"github.com/markelog/pilgrima/test/schema"
	testfixtures "gopkg.in/testfixtures.v2"
)

var (
	app *iris.Application
	fx  *testfixtures.Context
	db  *gorm.DB
)

func prepare() *iris.Application {
	if err := fx.Load(); err != nil {
		log.Fatal(err)
	}

	return app
}

func teardown() {
	db.Exec("TRUNCATE users CASCADE;")
	db.Exec("TRUNCATE projects CASCADE;")
	db.Exec("TRUNCATE branches CASCADE;")
	db.Exec("TRUNCATE commits CASCADE;")
	db.Exec("TRUNCATE reports CASCADE;")
	db.Exec("TRUNCATE tokens CASCADE;")
}
func TestMain(m *testing.M) {
	env.Up()

	app = routes.Up()
	db = database.Up()
	fx = fixtures.Up("fixtures", db)
	log := logger.Up()
	log.Out = ioutil.Discard

	token.Up(app, db, log)

	app.Build()

	os.Exit(m.Run())
}

func TestError(t *testing.T) {
	defer teardown()
	req := request.Up(app, t)

	token := req.POST("/token").
		WithHeader("Content-Type", "application/json").
		Expect().
		Status(http.StatusBadRequest)

	token.JSON().Schema(schema.Response)
}

func TestSuccess(t *testing.T) {
	defer teardown()
	prepare()
	req := request.Up(app, t)

	data := map[string]interface{}{
		"project": 1,
	}

	token := req.POST("/token").
		WithHeader("Content-Type", "application/json").
		WithJSON(data).
		Expect()

	token.JSON().Schema(schema.Response)
}
