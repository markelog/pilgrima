package root

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/markelog/pilgrima/database/models"
	"github.com/sirupsen/logrus"
)

var project models.Project

// Up root route
func Up(app *iris.Application, db *gorm.DB, logger *logrus.Logger) {
	app.Get("/", func(ctx iris.Context) {
		db.First(&project)
		ctx.HTML("<h1>" + project.Name + "</h1>")
	})
}
