package token_test

import (
	"log"
	"net/http"
	"testing"

	"github.com/joho/godotenv"
	"gopkg.in/gavv/httpexpect.v1"

	"github.com/markelog/pilgrima/application"
	"github.com/markelog/pilgrima/database"
)

func irisTester(t *testing.T) *httpexpect.Expect {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Panic(err)
	}

	var (
		app = application.Up(database.Up())
	)

	app.Build()

	return httpexpect.WithConfig(httpexpect.Config{
		BaseURL: "http://example.com",
		Client: &http.Client{
			Transport: httpexpect.NewBinder(app),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewRequireReporter(t),
	})
}

func TestError(t *testing.T) {
	request := irisTester(t)

	schema := `{
		"type": "object",
	    "properties": {
			"message": {"type": "string"},
			"payload": {"type": "object"},
			"status":  {"type": "string"}
	    },
	    "required": ["message", "status", "payload"]
	}`

	token := request.POST("/token").
		WithHeader("Content-Type", "application/json").
		Expect().
		Status(http.StatusBadRequest)

	token.JSON().Schema(schema)
}
