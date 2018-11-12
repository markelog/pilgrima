package report

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	controller "github.com/markelog/pilgrima/controllers/report"
	"github.com/sirupsen/logrus"
)

func setError(log *logrus.Logger, params *controller.CreateArgs, err error) {
	log.WithFields(logrus.Fields{
		"project": params.Project.Repository,
		"branch":  params.Project.Branch.Name,
	}).Error(err.Error())

	ctx.StatusCode(iris.StatusBadRequest)
	ctx.JSON(iris.Map{
		"status":  "failed",
		"message": "Can't create the report",
		"payload": iris.Map{},
	})
}

// Up report route
func Up(app *iris.Application, db *gorm.DB, log *logrus.Logger) {
	app.Post("/report", func(ctx iris.Context) {
		var params controller.CreateArgs
		err := ctx.ReadJSON(&params)

		if err != nil {
			setError(log, &params, err)
			return
		}

		ctrl := controller.New(db)
		err = ctrl.Create(&params)

		if err != nil {
			setError(log, &params, err)
			return
		}

		log.WithFields(logrus.Fields{
			"reports": args.Project.Branch.Commit.Report,
		}).Info("Reports created")

		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(iris.Map{
			"status":  "created",
			"message": "Yey!",
			"payload": iris.Map{},
		})
	})
}
