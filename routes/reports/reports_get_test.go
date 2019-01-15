package reports_test

import (
	"net/http"
	"testing"

	"github.com/markelog/pilgrima/test/request"
)

func TestGet(t *testing.T) {
	teardown()
	defer teardown()
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
						"first.a": map[string]interface{}{
							"gzip": 1,
							"size": 2,
						},
						"first.b": map[string]interface{}{
							"gzip": 3,
							"size": 4,
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
					"hash":    "16adf584c366f9626c6b799b69de41d0a11acef2",
					"author":  "Oleg Gaidarenko <markelog@gmail.com>",
					"message": "Sup",
					"report": map[string]interface{}{
						"first.a": map[string]interface{}{
							"gzip": 5,
							"size": 6,
						},
						"first.b": map[string]interface{}{
							"gzip": 7,
							"size": 8,
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
					"hash":    "aaff86b26b581f367ef099b4a2015b875ec2aa79",
					"author":  "Oleg Gaidarenko <markelog@gmail.com>",
					"message": "Work on the report route",
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
					"hash":    "2a5a7a2a60a36ab64546caaa10f10a39b14e37f7",
					"author":  "dependabot[bot] <support@dependabot.com>",
					"message": "Sup",
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

	req.POST("/reports").
		WithHeader("Content-Type", "application/json").
		WithJSON(first).
		Expect().
		Status(http.StatusOK)

	req.POST("/reports").
		WithHeader("Content-Type", "application/json").
		WithJSON(second).
		Expect().
		Status(http.StatusOK)

	req.POST("/reports").
		WithHeader("Content-Type", "application/json").
		WithJSON(third).
		Expect().
		Status(http.StatusOK)

	req.POST("/reports").
		WithHeader("Content-Type", "application/json").
		WithJSON(fourth).
		Expect().
		Status(http.StatusOK)

	response := req.GET("/reports").
		WithQuery("repository", "github.com/markelog/adit").
		WithQuery("branch", "master").
		WithHeader("Content-Type", "application/json").
		Expect().
		Status(http.StatusOK)

	payload := response.JSON().Object().Value("payload").Array()

	firstElement := payload.Element(0).Object()

	firstElement.Value("author").Equal("Oleg Gaidarenko \u003cmarkelog@gmail.com\u003e")
	firstElement.Value("hash").Equal("16adf584c366f9626c6b799b69de41d0a11acef2")
	firstElement.Value("message").Equal("Sup")

	sizes := firstElement.Value("sizes").Array()

	firstVal := sizes.Element(0).Object()
	secondVal := sizes.Element(1).Object()

	firstVal.Value("gzip").Equal(5)
	firstVal.Value("name").Equal("first.a")
	firstVal.Value("size").Equal(6)

	secondVal.Value("gzip").Equal(7)
	secondVal.Value("name").Equal("first.b")
	secondVal.Value("size").Equal(8)

	secondElement := payload.Element(1).Object()

	secondElement.Value("author").Equal("Oleg Gaidarenko \u003cmarkelog@gmail.com\u003e")
	secondElement.Value("hash").Equal("952b6fd9f671baa3719d680c508f828d12a893cd")
	secondElement.Value("message").Equal("Sup")

	sizes = secondElement.Value("sizes").Array()

	firstVal = sizes.Element(0).Object()
	secondVal = sizes.Element(1).Object()

	firstVal.Value("gzip").Equal(1)
	firstVal.Value("name").Equal("first.a")
	firstVal.Value("size").Equal(2)

	secondVal.Value("gzip").Equal(3)
	secondVal.Value("name").Equal("first.b")
	secondVal.Value("size").Equal(4)
}

// func TestGetNotFound(t *testing.T) {
// 	defer teardown()
// 	req := request.Up(app, t)

// 	response := req.GET("/reports").
// 		WithQuery("repository", "github.com/markelog/adit").
// 		WithQuery("branch", "test").
// 		WithHeader("Content-Type", "application/json").
// 		Expect().
// 		Status(http.StatusNotFound)

// 	json := response.JSON()

// 	json.Schema(schema.Response)
// 	json.Object().Value("payload").Object().Empty()
// 	json.Object().Value("message").Equal("Not found")
// 	json.Object().Value("status").Equal("failed")
// }
