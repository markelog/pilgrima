package token_test

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/markelog/pilgrima/database"
	"github.com/markelog/pilgrima/database/models"
	"github.com/markelog/pilgrima/logger"
	"github.com/markelog/pilgrima/routes/token"
	"github.com/markelog/pilgrima/test/env"
	"github.com/markelog/pilgrima/test/fixtures"
	"github.com/markelog/pilgrima/test/request"
	"github.com/markelog/pilgrima/test/routes"
	"github.com/markelog/pilgrima/test/schema"
	testfixtures "gopkg.in/testfixtures.v2"
)

var (
	app *iris.Application
	fx  *testfixtures.Context
	db  *gorm.DB
)

func prepare() *iris.Application {
	if err := fx.Load(); err != nil {
		log.Fatal(err)
	}

	return app
}

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
	fx = fixtures.Up("fixtures", db)
	log := logger.Up()
	log.Out = ioutil.Discard

	token.Up(app, db, log)

	app.Build()
	// teardown()

	os.Exit(m.Run())
}

func TestError(t *testing.T) {
	// defer teardown()
	req := request.Up(app, t)

	token := req.POST("/token").
		WithHeader("Content-Type", "routes/json").
		Expect().
		Status(http.StatusBadRequest)

	token.JSON().Schema(schema.Response)
}

func TestSuccess(t *testing.T) {
	// defer teardown()
	// prepare()
	// req := request.Up(app, t)

	// data := map[string]interface{}{
	// 	"project": 1,
	// }

	// token := req.POST("/token").
	// 	WithHeader("Content-Type", "routes/json").
	// 	WithJSON(data).
	// 	Expect().
	// 	Status(http.StatusOK)

	// token.JSON().Schema(schema.Response)
}
