package report

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	controller "github.com/markelog/pilgrima/controllers/report"
	"github.com/sirupsen/logrus"
)

// Up report route
func Up(app *iris.Application, db *gorm.DB, log *logrus.Logger) {
	app.Post("/report", func(ctx iris.Context) {
		var params controller.CreateArgs
		ctx.ReadJSON(&params.Project)

		ctrl := controller.New(db)
		result := ctrl.Create(&params)

		spew.Dump(result.Error)
	})
}
