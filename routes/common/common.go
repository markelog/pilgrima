package common

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/sirupsen/logrus"
)

// Up project route
func Up(app *iris.Application, db *gorm.DB, log *logrus.Logger) {
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		log.WithFields(logrus.Fields{
			"url":    ctx.Path(),
			"method": ctx.Method(),
		}).Error("Returned 404 error")

		ctx.StatusCode(iris.StatusNotFound)
		ctx.JSON(iris.Map{
			"status":  "error",
			"message": "Can't find this route",
			"payload": iris.Map{},
		})
	})
}
