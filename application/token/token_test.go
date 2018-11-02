package token_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/markelog/pilgrima/application/token"
	"github.com/markelog/pilgrima/database"
	"github.com/markelog/pilgrima/test/application"
	"github.com/markelog/pilgrima/test/env"
	"github.com/markelog/pilgrima/test/fixtures"
	"github.com/markelog/pilgrima/test/request"
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
func TestMain(m *testing.M) {
	env.Up()

	app = application.Up()
	db = database.Up()
	fx = fixtures.Up("fixtures", db)

	token.Up(app, db)

	app.Build()

	os.Exit(m.Run())
}

func TestError(t *testing.T) {
	req := request.Up(app, t)

	token := req.POST("/token").
		WithHeader("Content-Type", "application/json").
		Expect().
		Status(http.StatusBadRequest)

	token.JSON().Schema(schema.Response)
}

func TestSuccess(t *testing.T) {
	prepare()
	req := request.Up(app, t)

	data := map[string]interface{}{
		"project": 1,
	}

	token := req.POST("/token").
		WithHeader("Content-Type", "application/json").
		WithJSON(data).
		Expect().
		Status(http.StatusOK)

	token.JSON().Schema(schema.Response)
}
