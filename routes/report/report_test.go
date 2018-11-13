package report_test

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/markelog/pilgrima/database"
	"github.com/markelog/pilgrima/database/models"
	"github.com/markelog/pilgrima/logger"
	"github.com/markelog/pilgrima/routes/report"
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
	db.Unscoped().Delete(&models.User{})
	db.Unscoped().Delete(&models.Project{})
	db.Unscoped().Delete(&models.Branch{})
	db.Unscoped().Delete(&models.Commit{})
	db.Unscoped().Delete(&models.Report{})
	db.Unscoped().Delete(&models.Token{})
}

func TestMain(m *testing.M) {
	env.Up()

	app = routes.Up()
	db = database.Up()
	log := logger.Up()
	log.Out = ioutil.Discard

	report.Up(app, db, log)

	app.Build()

	os.Exit(m.Run())
}

func TestError(t *testing.T) {
	defer teardown()
	req := request.Up(app, t)

	data := map[string]interface{}{
		"project": &map[string]interface{}{
			"repository": "https://github.com/markelog/adit",
			"branch": map[string]interface{}{
				"name": "master",
				"commit": map[string]interface{}{
					"hash":      "952b6fd9f671baa3719d680c508f828d12a893cd",
					"committer": "Oleg Gaidarenko <markelog@gmail.com>",
					"message":   "Sup",
					"report": []map[string]interface{}{
						map[string]interface{}{
							"name": "test",
							"size": "nope!",
							"gzip": 123,
						},
						map[string]interface{}{
							"name": "super",
							"size": 321,
							"gzip": 123,
						},
					},
				},
			},
		},
	}

	response := req.POST("/report").
		WithHeader("Content-Type", "application/json").
		WithJSON(data).
		Expect().
		Status(http.StatusBadRequest)

	json := response.JSON()

	json.Schema(schema.Response)

	spew.Dump(json)

	json.Object().
		Value("message").Equal("Can't create the report")

	json.Object().
		Value("status").Equal("failed")
}

func TestSuccess(t *testing.T) {
	defer teardown()
	req := request.Up(app, t)

	data := map[string]interface{}{
		"project": &map[string]interface{}{
			"repository": "https://github.com/markelog/adit",
			"branch": map[string]interface{}{
				"name": "master",
				"commit": map[string]interface{}{
					"hash":      "952b6fd9f671baa3719d680c508f828d12a893cd",
					"committer": "Oleg Gaidarenko <markelog@gmail.com>",
					"message":   "Sup",
					"report": []map[string]interface{}{
						map[string]interface{}{
							"name": "test",
							"size": 9999,
							"gzip": 123,
						},
						map[string]interface{}{
							"name": "super",
							"size": 321,
							"gzip": 123,
						},
					},
				},
			},
		},
	}

	response := req.POST("/report").
		WithHeader("Content-Type", "application/json").
		WithJSON(data).
		Expect().
		Status(http.StatusOK)

	response.JSON().Schema(schema.Response)
}

func TestSuccessForSecondTime(t *testing.T) {
	defer teardown()
	req := request.Up(app, t)

	data := map[string]interface{}{
		"project": &map[string]interface{}{
			"repository": "https://github.com/markelog/adit",
			"branch": map[string]interface{}{
				"name": "master",
				"commit": map[string]interface{}{
					"hash":      "952b6fd9f671baa3719d680c508f828d12a893cd",
					"committer": "Oleg Gaidarenko <markelog@gmail.com>",
					"message":   "Sup",
					"report": []map[string]interface{}{
						map[string]interface{}{
							"name": "test",
							"size": 9999,
							"gzip": 123,
						},
						map[string]interface{}{
							"name": "super",
							"size": 321,
							"gzip": 123,
						},
					},
				},
			},
		},
	}

	req.POST("/report").
		WithHeader("Content-Type", "application/json").
		WithJSON(data)

	response := req.POST("/report").
		WithHeader("Content-Type", "application/json").
		WithJSON(data).
		Expect().
		Status(http.StatusOK)

	response.JSON().Schema(schema.Response)
}
