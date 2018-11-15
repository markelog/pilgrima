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
func TestGetLast(t *testing.T) {
	defer teardown()
	req := request.Up(app, t)

	first := map[string]interface{}{
		"project": &map[string]interface{}{
			"repository": "github.com/markelog/adit",
			"branch": map[string]interface{}{
				"name": "master",
				"commit": map[string]interface{}{
					"hash":      "952b6fd9f671baa3719d680c508f828d12a893cd",
					"committer": "Oleg Gaidarenko <markelog@gmail.com>",
					"message":   "Sup",
					"report": []map[string]interface{}{
						map[string]interface{}{
							"name": "first.a",
							"size": 9999,
							"gzip": 123,
						},
						map[string]interface{}{
							"name": "first.b",
							"size": 321,
							"gzip": 123,
						},
					},
				},
			},
		},
	}

	second := map[string]interface{}{
		"project": &map[string]interface{}{
			"repository": "github.com/markelog/adit",
			"branch": map[string]interface{}{
				"name": "master",
				"commit": map[string]interface{}{
					"hash":      "16adf584c366f9626c6b799b69de41d0a11acef2",
					"committer": "Oleg Gaidarenko <markelog@gmail.com>",
					"message":   "Sup",
					"report": []map[string]interface{}{
						map[string]interface{}{
							"name": "second.a",
							"size": 9999,
							"gzip": 123,
						},
						map[string]interface{}{
							"name": "second.b",
							"size": 321,
							"gzip": 123,
						},
					},
				},
			},
		},
	}

	third := map[string]interface{}{
		"project": &map[string]interface{}{
			"repository": "github.com/oleg-koval/ya-skeleton",
			"branch": map[string]interface{}{
				"name": "WIP",
				"commit": map[string]interface{}{
					"hash":      "aaff86b26b581f367ef099b4a2015b875ec2aa79",
					"committer": "Oleg Gaidarenko <markelog@gmail.com>",
					"message":   "Work on the report route",
					"report": []map[string]interface{}{
						map[string]interface{}{
							"name": "third.a",
							"size": 9999,
							"gzip": 123,
						},
						map[string]interface{}{
							"name": "third.b",
							"size": 321,
							"gzip": 123,
						},
					},
				},
			},
		},
	}

	fourth := map[string]interface{}{
		"project": &map[string]interface{}{
			"repository": "github.com/oleg-koval/ya-skeleton",
			"branch": map[string]interface{}{
				"name": "master",
				"commit": map[string]interface{}{
					"hash":      "2a5a7a2a60a36ab64546caaa10f10a39b14e37f7",
					"committer": "dependabot[bot] <support@dependabot.com>",
					"message":   "Sup",
					"report": []map[string]interface{}{
						map[string]interface{}{
							"name": "fourth.a",
							"size": 9999,
							"gzip": 123,
						},
						map[string]interface{}{
							"name": "fourth.b",
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
		WithJSON(first).
		Expect().
		Status(http.StatusOK)

	req.POST("/report").
		WithHeader("Content-Type", "application/json").
		WithJSON(second).
		Expect().
		Status(http.StatusOK)

	req.POST("/report").
		WithHeader("Content-Type", "application/json").
		WithJSON(third).
		Expect().
		Status(http.StatusOK)

	req.POST("/report").
		WithHeader("Content-Type", "application/json").
		WithJSON(fourth).
		Expect().
		Status(http.StatusOK)

	response := req.GET("report/last?repository=github.com/markelog/adit&branch=master").
		WithQuery("repository", "github.com/markelog/adit").
		WithQuery("branch", "master").
		WithHeader("Content-Type", "application/json").
		Expect()
	// Status(http.StatusOK)

	spew.Dump(response.JSON())
}
