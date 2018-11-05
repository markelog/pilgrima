package project_test

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/markelog/pilgrima/database"
	"github.com/markelog/pilgrima/database/models"
	"github.com/markelog/pilgrima/logger"
	"github.com/markelog/pilgrima/routes/project"
	"github.com/markelog/pilgrima/test/env"
	"github.com/markelog/pilgrima/test/request"
	"github.com/markelog/pilgrima/test/routes"
	"github.com/markelog/pilgrima/test/schema"
)

var (
	app *iris.Application
	db  *gorm.DB
)

func teardown() {
	db.Delete(&models.Project{})
	db.Delete(&models.Token{})
}

func TestMain(m *testing.M) {
	env.Up()

	app = routes.Up()
	db = database.Up()
	log := logger.Up()
	log.Out = ioutil.Discard

	project.Up(app, db, log)

	app.Build()

	os.Exit(m.Run())
}

func TestAbsenceOfARepository(t *testing.T) {
	defer teardown()
	req := request.Up(app, t)

	data := map[string]interface{}{
		"name": "test",
	}

	token := req.POST("/project").
		WithHeader("Content-Type", "routes/json").
		WithJSON(data).
		Expect().
		Status(http.StatusBadRequest)

	json := token.JSON()

	json.Schema(schema.Response)

	json.Object().
		Value("payload").Object().
		Value("errors").Array().
		Elements("repository: Does not match format 'uri'")
}

func TestAbsenceOfAName(t *testing.T) {
	defer teardown()
	req := request.Up(app, t)

	data := map[string]interface{}{
		"repository": "https://github.com/markelog/pilgrima",
	}

	token := req.POST("/project").
		WithHeader("Content-Type", "routes/json").
		WithJSON(data).
		Expect().
		Status(http.StatusBadRequest)

	json := token.JSON()

	json.Schema(schema.Response)

	json.Object().
		Value("payload").Object().
		Value("errors").Array().
		Elements("name: String length must be greater than or equal to 1")
}

func TestAbsence(t *testing.T) {
	defer teardown()
	req := request.Up(app, t)

	token := req.POST("/project").
		WithHeader("Content-Type", "routes/json").
		Expect().
		Status(http.StatusBadRequest)

	json := token.JSON()

	json.Schema(schema.Response)

	json.Object().
		Value("payload").Object().
		Value("errors").Array().
		Elements(
			"name: String length must be greater than or equal to 1",
			"repository: Does not match format 'uri'",
		)
}

func TestSuccess(t *testing.T) {
	defer teardown()
	req := request.Up(app, t)

	data := map[string]interface{}{
		"name":       "yo",
		"repository": "http://github.com/markelog/pilgrima",
	}

	token := req.POST("/project").
		WithHeader("Content-Type", "routes/json").
		WithJSON(data).
		Expect().
		Status(http.StatusOK)

	token.JSON().Schema(schema.Response)
}
