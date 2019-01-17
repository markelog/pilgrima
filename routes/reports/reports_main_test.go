package reports_test

import (
	"net/http"
	"testing"

	"github.com/markelog/pilgrima/test/request"
	"github.com/markelog/pilgrima/test/schema"
)

func TestError(t *testing.T) {
	teardown()
	defer teardown()
	req := request.Up(app, t)

	data := map[string]interface{}{
		"project": &map[string]interface{}{
			"repository": "github.com/markelog/adit",
			"branch": map[string]interface{}{
				"name": "master",
				"commit": map[string]interface{}{
					"hash":    "952b6fd9f671baa3719d680c508f828d12a893cd",
					"author":  "Oleg Gaidarenko <markelog@gmail.com>",
					"message": "Sup",
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

	response := req.POST("/reports").
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
	teardown()
	defer teardown()
	req := request.Up(app, t)

	data := map[string]interface{}{
		"project": &map[string]interface{}{
			"repository": "",
			"branch": map[string]interface{}{
				"name": "master",
				"commit": map[string]interface{}{
					"hash":    "952b6fd9f671baa3719d680c508f828d12a893cd",
					"author":  "Oleg Gaidarenko <markelog@gmail.com>",
					"message": "Sup",
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

	response := req.POST("/reports").
		WithHeader("Content-Type", "application/json").
		WithJSON(data).
		Expect().
		Status(http.StatusBadRequest)

	json := response.JSON()

	json.Schema(schema.Response)

	json.Object().Value("message").Equal("repository: repository is required")
}

func TestSuccess(t *testing.T) {
	teardown()
	defer teardown()
	req := request.Up(app, t)

	data := map[string]interface{}{
		"project": &map[string]interface{}{
			"repository": "github.com/markelog/adit",
			"branch": map[string]interface{}{
				"name": "master",
				"commit": map[string]interface{}{
					"hash":    "952b6fd9f671baa3719d680c508f828d12a893cd",
					"author":  "Oleg Gaidarenko <markelog@gmail.com>",
					"message": "Sup",
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

	response := req.POST("/reports").
		WithHeader("Content-Type", "application/json").
		WithJSON(data).
		Expect().
		Status(http.StatusOK)

	response.JSON().Schema(schema.Response)
}

func TestSuccessForSecondTime(t *testing.T) {
	teardown()
	// defer teardown()
	req := request.Up(app, t)

	first := map[string]interface{}{
		"project": map[string]interface{}{
			"repository": "github.com/markelog/adit",
			"branch": map[string]interface{}{
				"name": "master",
				"commit": map[string]interface{}{
					"hash":    "952b6fd9f671baa3719d680c508f828d12a893cd",
					"author":  "Oleg Gaidarenko <markelog@gmail.com>",
					"message": "Sup",
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

	second := map[string]interface{}{
		"project": map[string]interface{}{
			"repository": "github.com/markelog/adit",
			"branch": map[string]interface{}{
				"name": "master",
				"commit": map[string]interface{}{
					"hash":    "952b6fd9f671baa3719d680c508f828d12a893cd",
					"author":  "Killa Gorilla <killa@gorilla.com>",
					"message": "Sup",
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

	req.POST("/reports").
		WithHeader("Content-Type", "application/json").
		WithJSON(first).
		Expect().
		Status(http.StatusOK)

	response := req.POST("/reports").
		WithHeader("Content-Type", "application/json").
		WithJSON(second).
		Expect().
		Status(http.StatusOK)

	response.JSON().Schema(schema.Response)

	response = req.GET("/reports").
		WithQuery("repository", "github.com/markelog/adit").
		WithQuery("branch", "master").
		WithHeader("Content-Type", "application/json").
		Expect()

	response.JSON().Object().Value("payload").Array().Element(0).
		Object().Value("author").Equal("Killa Gorilla <killa@gorilla.com>")
}
