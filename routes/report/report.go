package report

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	controller "github.com/markelog/pilgrima/controllers/report"
	"github.com/sirupsen/logrus"
)

func setPostError(log *logrus.Logger, params *controller.CreateArgs, ctx iris.Context, err error) {
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

func setLastError(log *logrus.Logger, params *controller.LastArgs, ctx iris.Context, err error) {
	log.WithFields(logrus.Fields{
		"project": params.Repository,
		"branch":  params.Branch,
	}).Error(err.Error())

	ctx.StatusCode(iris.StatusBadRequest)
	ctx.JSON(iris.Map{
		"status":  "failed",
		"message": "Can't find those reports",
		"payload": iris.Map{},
	})
}

// Up report route
func Up(app *iris.Application, db *gorm.DB, log *logrus.Logger) {
	ctrl := controller.New(db)

	app.Post("/report", func(ctx iris.Context) {
		var params controller.CreateArgs
		err := ctx.ReadJSON(&params)

		if err != nil {
			setPostError(log, &params, ctx, err)
			return
		}

		err = ctrl.Create(&params)

		if err != nil {
			setPostError(log, &params, ctx, err)
			return
		}

		log.WithFields(logrus.Fields{
			"reports": params.Project.Branch.Commit.Report,
		}).Info("Reports created")

		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(iris.Map{
			"status":  "created",
			"message": "Yey!",
			"payload": iris.Map{},
		})
	})

	app.Get("/report/last", func(ctx iris.Context) {
		URLparams := ctx.URLParams()

		params := controller.LastArgs{
			Repository: URLparams["repository"],
			Branch:     URLparams["branch"],
		}

		reports, err := ctrl.Last(&params)
		if err != nil {
			setLastError(log, &params, ctx, err)
			return
		}

		payload, err := json.Marshal(reports)
		if err != nil {
			setLastError(log, &params, ctx, err)
			return
		}

		log.WithFields(logrus.Fields{
			"repository": URLparams["repository"],
			"branch":     URLparams["branch"],
		}).Info("Last reports will be returned")

		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(iris.Map{
			"status":  "success",
			"message": "Found",
			"payload": string(payload),
		})
	})
}
