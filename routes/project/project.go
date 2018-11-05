package token

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	controller "github.com/markelog/pilgrima/controllers/project"
	"github.com/markelog/pilgrima/controllers/token"
	"github.com/sirupsen/logrus"
)

type postProject struct {
	Project    string `json:"name"`
	Repository string `json:"repository"`
}

// Up project route
func Up(app *iris.Application, db *gorm.DB, log *logrus.Logger) {
	app.Post("/project", func(ctx iris.Context) {
		var params postProject
		ctx.ReadJSON(&params)

		ctrl := controller.New(
			params.Project,
			params.Repository,
			db,
		)

		tx := db.Begin()
		result, value := ctrl.Create()

		if result.Error != nil {
			log.WithFields(logrus.Fields{
				"project":    params.Project,
				"repository": params.Repository,
			}).Error("Can't create the project")

			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{
				"status":  "failed",
				"message": "Can't create the project",
				"payload": iris.Map{},
			})

			return
		}

		tokenCtrl := token.New(db)
		tokenResult, _ := tokenCtrl.Create(value.ID)

		if tokenResult.Error != nil {
			tx.Rollback()

			log.WithFields(logrus.Fields{
				"project":    params.Project,
				"repository": params.Repository,
			}).Error("Can't create the token")

			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{
				"status":  "failed",
				"message": "Can't create the token",
				"payload": iris.Map{},
			})

			return
		}

		tx.Commit()

		log.WithFields(logrus.Fields{
			"project": params.Project,
		}).Info("Project created")

		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(iris.Map{
			"status":  "created",
			"message": "Yey!",
			"payload": iris.Map{
				"project": ctrl.Name,
			},
		})
	})
}
