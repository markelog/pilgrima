package token

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	controller "github.com/markelog/pilgrima/controllers/token"
	"github.com/sirupsen/logrus"
)

type postProject struct {
	Project uint `json:"project"`
}

// Up token route
func Up(app *iris.Application, db *gorm.DB, log *logrus.Logger) {
	app.Post("/token", func(ctx iris.Context) {
		var params postProject
		ctx.ReadJSON(&params)

		result, value := controller.New(db).Create(params.Project)

		if result.RecordNotFound() {
			log.WithFields(logrus.Fields{
				"project": params.Project,
			}).Error("Can't find this project")

			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{
				"status":  "failed",
				"message": "Can't find this project",
				"payload": iris.Map{},
			})
			return
		}

		if result.Error != nil {
			log.WithFields(logrus.Fields{
				"project": params.Project,
				"error":   result.Error,
			}).Error("Couldn't create the token")

			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{
				"status":  "failed",
				"message": "Something wen't wrong",
				"payload": iris.Map{},
			})
			return
		}

		log.WithFields(logrus.Fields{
			"project": params.Project,
			"token":   value.Token,
		}).Info("Token created")

		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(iris.Map{
			"status":  "created",
			"message": "Yey!",
			"payload": iris.Map{
				"token": value.Token,
			},
		})
	})
}
