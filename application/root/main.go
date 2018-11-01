package root

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/markelog/pilgrima/database/models"
)

var project models.Project

// Set root route
func Set(app *iris.Application, db *gorm.DB) {
	app.Get("/", func(ctx iris.Context) {
		db.First(&project)
		ctx.HTML("<h1>" + project.Name + "</h1>")
	})
}
