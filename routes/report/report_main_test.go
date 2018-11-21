package report_test

import (
	"net/http"
	"testing"

	"github.com/markelog/pilgrima/test/request"
	"github.com/markelog/pilgrima/test/schema"
)

func TestError(t *testing.T) {
	defer teardown()
	req := request.Up(app, t)

	data := map[string]interface{}{
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
			"repository": "github.com/markelog/adit",
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
			"repository": "github.com/markelog/adit",
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
