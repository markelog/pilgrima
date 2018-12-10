package common_test

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/markelog/pilgrima/database"
	"github.com/markelog/pilgrima/logger"
	"github.com/markelog/pilgrima/routes/common"
	"github.com/markelog/pilgrima/test/env"
	"github.com/markelog/pilgrima/test/request"
	"github.com/markelog/pilgrima/test/routes"
	"github.com/markelog/pilgrima/test/schema"
)

var (
	app *iris.Application
	db  *gorm.DB
)

func TestMain(m *testing.M) {
	env.Up()

	app = routes.Up()
	db = database.Up()
	log := logger.Up()
	log.Out = ioutil.Discard

	common.Up(app, db, log)

	app.Build()

	os.Exit(m.Run())
}

func TestPOST404(t *testing.T) {
	req := request.Up(app, t)

	response := req.POST("/not-found").
		WithHeader("Content-Type", "routes/json").
		Expect().
		Status(http.StatusNotFound)

	json := response.JSON()

	json.Schema(schema.Response)

	json.Object().Value("status").Equal("error")
	json.Object().Value("message").Equal("Can't this route")
	json.Object().Value("payload").Object().Empty()
}
