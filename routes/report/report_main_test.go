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
					"report": map[string]interface{}{
						"test": map[string]interface{}{
							"size": "nope!",
							"gzip": 123,
						},
						"super": map[string]interface{}{
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
		Value("message").Equal("json: cannot unmarshal string into Go struct field .size of type uint")

	json.Object().
		Value("status").Equal("failed")
}

func TestEmptyProject(t *testing.T) {
	defer teardown()
	req := request.Up(app, t)

	data := map[string]interface{}{
		"project": &map[string]interface{}{
			"repository": "",
			"branch": map[string]interface{}{
				"name": "master",
				"commit": map[string]interface{}{
					"hash":      "952b6fd9f671baa3719d680c508f828d12a893cd",
					"committer": "Oleg Gaidarenko <markelog@gmail.com>",
					"message":   "Sup",
					"report": map[string]interface{}{
						"test": map[string]interface{}{
							"size": 9999,
							"gzip": 123,
						},
						"super": map[string]interface{}{
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

	json.Object().Value("message").Equal("repository: repository is required")
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
					"report": map[string]interface{}{
						"test": map[string]interface{}{
							"size": 9999,
							"gzip": 123,
						},
						"super": map[string]interface{}{
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
		"project": map[string]interface{}{
			"repository": "github.com/markelog/adit",
			"branch": map[string]interface{}{
				"name": "master",
				"commit": map[string]interface{}{
					"hash":      "952b6fd9f671baa3719d680c508f828d12a893cd",
					"committer": "Oleg Gaidarenko <markelog@gmail.com>",
					"message":   "Sup",
					"report": map[string]interface{}{
						"test": map[string]interface{}{
							"size": 9999,
							"gzip": 123,
						},
						"super": map[string]interface{}{
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
