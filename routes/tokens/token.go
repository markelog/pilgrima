package tokens

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
	controller "github.com/markelog/pilgrima/controllers/token"
	"github.com/sirupsen/logrus"
)

type postProject struct {
	Project uint `json:"project"`
}

// Up token route
func Up(app *iris.Application, db *gorm.DB, log *logrus.Logger) {
	ctrl := controller.New(db)

	app.Post("/tokens", func(ctx iris.Context) {
		var params postProject
		ctx.ReadJSON(&params)

		result, err := ctrl.Create(params.Project)

		if err != nil && err.Error() == "record not found" {
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

		if err != nil {
			log.WithFields(logrus.Fields{
				"project": params.Project,
				"error":   err,
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
			"token":   result.Token,
		}).Info("Token created")

		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(iris.Map{
			"status":  "created",
			"message": "Yey!",
			"payload": iris.Map{
				"token": result.Token,
			},
		})
	})
}
