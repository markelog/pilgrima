package app

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/markelog/pilgrima/models"
)

var project models.Project

// Root route
func Root(app *iris.Application, db *gorm.DB) {
	app.Get("/", func(ctx iris.Context) {
		db.First(&project)
		ctx.HTML("<h1>" + project.Name + "</h1>")
	})
}
