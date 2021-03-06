package projects_test

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
	"github.com/markelog/pilgrima/database"
	"github.com/markelog/pilgrima/logger"
	"github.com/markelog/pilgrima/routes/projects"
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
	db.Raw("TRUNCATE users CASCADE;").Row()
	db.Raw("TRUNCATE projects CASCADE;").Row()
	db.Raw("TRUNCATE tokens CASCADE;").Row()
}

func TestMain(m *testing.M) {
	env.Up()

	app = routes.Up()
	db = database.Up()
	log := logger.Up()
	log.Out = ioutil.Discard

	projects.Up(app, db, log)

	app.Build()

	os.Exit(m.Run())
}

func TestAbsenceOfARepository(t *testing.T) {
	defer teardown()
	req := request.Up(app, t)

	data := map[string]interface{}{
		"name": "test",
	}

	token := req.POST("/projects").
		WithHeader("Content-Type", "application/json").
		WithJSON(data).
		Expect().
		Status(http.StatusBadRequest)

	json := token.JSON()

	json.Schema(schema.Response)

	json.Object().
		Value("payload").Object().
		Value("errors").Array().
		Elements("repository: String length must be greater than or equal to 1")
}

func TestAbsenceOfAName(t *testing.T) {
	defer teardown()
	req := request.Up(app, t)

	data := map[string]interface{}{
		"repository": "https://github.com/markelog/pilgrima",
	}

	token := req.POST("/projects").
		WithHeader("Content-Type", "application/json").
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

	token := req.POST("/projects").
		WithHeader("Content-Type", "application/json").
		Expect().
		Status(http.StatusBadRequest)

	json := token.JSON()

	json.Schema(schema.Response)

	json.Object().
		Value("payload").Object().
		Value("errors").Array().
		Contains(
			"name: String length must be greater than or equal to 1",
		)
}

func TestSuccess(t *testing.T) {
	defer teardown()
	req := request.Up(app, t)

	data := map[string]interface{}{
		"name":       "yo",
		"repository": "github.com/markelog/pilgrima",
	}

	project := req.POST("/projects").
		WithHeader("Content-Type", "application/json").
		WithJSON(data).
		Expect().
		Status(http.StatusOK)

	project.JSON().Schema(schema.Response)
}

func TestList(t *testing.T) {
	defer teardown()
	teardown()
	req := request.Up(app, t)

	data := map[string]interface{}{
		"name":       "yo",
		"repository": "github.com/markelog/pilgrima",
	}

	req.POST("/projects").
		WithHeader("Content-Type", "application/json").
		WithJSON(data).
		Expect().
		Status(http.StatusOK)

	result := req.GET("/projects").
		Expect().
		Status(http.StatusOK).
		JSON()

	result.Schema(schema.Response)

	element := result.Object().Value("payload").Array().
		Element(0).Object()

	element.Value("name").Equal("yo")
	element.Value("repository").Equal("github.com/markelog/pilgrima")
}

func TestAbsentList(t *testing.T) {
	defer teardown()
	teardown()
	req := request.Up(app, t)

	response := req.GET("/projects").
		Expect().
		Status(http.StatusNotFound)

	json := response.JSON()

	json.Schema(schema.Response)

	json.Schema(schema.Response)
	json.Object().Value("payload").Object().Empty()
	json.Object().Value("message").Equal("Not found")
	json.Object().Value("status").Equal("failed")
}
