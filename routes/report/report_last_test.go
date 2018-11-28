package report_test

import (
	"net/http"
	"testing"

	"github.com/markelog/pilgrima/test/request"
	"github.com/markelog/pilgrima/test/schema"
)

func TestGetLast(t *testing.T) {
	defer teardown()
	req := request.Up(app, t)

	first := map[string]interface{}{
		"project": map[string]interface{}{
			"repository": "github.com/markelog/adit",
			"branch": map[string]interface{}{
				"name": "master",
				"commit": map[string]interface{}{
					"hash":      "952b6fd9f671baa3719d680c508f828d12a893cd",
					"committer": "Oleg Gaidarenko <markelog@gmail.com>",
					"message":   "Sup",
					"report": map[string]interface{}{
						"first.a": map[string]interface{}{
							"size": 2,
							"gzip": 1,
						},
						"first.b": map[string]interface{}{
							"size": 4,
							"gzip": 3,
						},
					},
				},
			},
		},
	}

	second := map[string]interface{}{
		"project": map[string]interface{}{
			"repository": "github.com/markelog/adit",
			"branch": map[string]interface{}{
				"name": "master",
				"commit": map[string]interface{}{
					"hash":      "16adf584c366f9626c6b799b69de41d0a11acef2",
					"committer": "Oleg Gaidarenko <markelog@gmail.com>",
					"message":   "Sup",
					"report": map[string]interface{}{
						"first.a": map[string]interface{}{
							"size": 6,
							"gzip": 5,
						},
						"first.b": map[string]interface{}{
							"size": 8,
							"gzip": 7,
						},
					},
				},
			},
		},
	}

	third := map[string]interface{}{
		"project": map[string]interface{}{
			"repository": "github.com/oleg-koval/ya-skeleton",
			"branch": map[string]interface{}{
				"name": "WIP",
				"commit": map[string]interface{}{
					"hash":      "aaff86b26b581f367ef099b4a2015b875ec2aa79",
					"committer": "Oleg Gaidarenko <markelog@gmail.com>",
					"message":   "Work on the report route",
					"report": map[string]interface{}{
						"third.a": map[string]interface{}{
							"size": 9999,
							"gzip": 123,
						},
						"third.b": map[string]interface{}{
							"size": 321,
							"gzip": 123,
						},
					},
				},
			},
		},
	}

	fourth := map[string]interface{}{
		"project": map[string]interface{}{
			"repository": "github.com/oleg-koval/ya-skeleton",
			"branch": map[string]interface{}{
				"name": "master",
				"commit": map[string]interface{}{
					"hash":      "2a5a7a2a60a36ab64546caaa10f10a39b14e37f7",
					"committer": "dependabot[bot] <support@dependabot.com>",
					"message":   "Sup",
					"report": map[string]interface{}{
						"fourth.a": map[string]interface{}{
							"size": 9999,
							"gzip": 123,
						},
						"fourth.b": map[string]interface{}{
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

	response := req.GET("/report/last").
		WithQuery("repository", "github.com/markelog/adit").
		WithQuery("branch", "master").
		WithHeader("Content-Type", "application/json").
		Expect().
		Status(http.StatusOK)

	json := response.JSON()
	object := json.Object().Value("payload").Object()

	firstVal := object.Value("first.a").Object()
	secondVal := object.Value("first.b").Object()

	firstVal.Value("size").Equal(6)
	firstVal.Value("gzip").Equal(5)

	secondVal.Value("size").Equal(8)
	secondVal.Value("gzip").Equal(7)
}

func TestNotFound(t *testing.T) {
	defer teardown()
	req := request.Up(app, t)

	response := req.GET("/report/last").
		WithQuery("repository", "github.com/markelog/adit").
		WithQuery("branch", "master").
		WithHeader("Content-Type", "application/json").
		Expect().
		Status(http.StatusNotFound)

	json := response.JSON()

	json.Schema(schema.Response)
	json.Object().Value("payload").Object().Empty()
}
