package reports

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	controller "github.com/markelog/pilgrima/controllers/report"
	"github.com/sirupsen/logrus"
)

func setPostError(log *logrus.Logger, params *controller.CreateArgs, ctx iris.Context, err error) {
	errorString := err.Error()

	log.WithFields(logrus.Fields{
		"project": params.Project.Repository,
		"branch":  params.Project.Branch.Name,
	}).Error(errorString)

	ctx.StatusCode(iris.StatusBadRequest)
	ctx.JSON(iris.Map{
		"status":  "failed",
		"message": errorString,
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
		"message": "Can't find that report",
		"payload": iris.Map{},
	})
}

func setGetError(log *logrus.Logger, params *controller.GetArgs, ctx iris.Context, err error) {
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

	app.Post("/reports", func(ctx iris.Context) {
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
			"report": params.Project.Branch.Commit.Report,
		}).Info("Report created")

		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(iris.Map{
			"status":  "created",
			"message": "Yey!",
			"payload": iris.Map{},
		})
	})

	app.Get("/reports/last", func(ctx iris.Context) {
		URLparams := ctx.URLParams()

		params := controller.LastArgs{
			Repository: URLparams["repository"],
			Branch:     URLparams["branch"],
		}

		report, err := ctrl.Last(&params)
		if err != nil {
			setLastError(log, &params, ctx, err)
			return
		}

		if len(report) == 0 {
			log.WithFields(logrus.Fields{
				"repository": URLparams["repository"],
				"branch":     URLparams["branch"],
			}).Info("Not found")

			ctx.StatusCode(iris.StatusNotFound)
			ctx.JSON(iris.Map{
				"status":  "failed",
				"message": "Not found",
				"payload": &controller.LastResult{},
			})
			return
		}

		log.WithFields(logrus.Fields{
			"repository": URLparams["repository"],
			"branch":     URLparams["branch"],
		}).Info("Last report will be returned")

		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(iris.Map{
			"status":  "success",
			"message": "Found",
			"payload": report,
		})
	})

	app.Get("/reports", func(ctx iris.Context) {
		URLparams := ctx.URLParams()

		params := controller.GetArgs{
			Repository: URLparams["repository"],
			Branch:     URLparams["branch"],
		}

		reports, err := ctrl.Get(&params)

		if err != nil {
			setGetError(log, &params, ctx, err)
			return
		}

		if len(reports) == 0 {
			log.WithFields(logrus.Fields{
				"repository": URLparams["repository"],
				"branch":     URLparams["branch"],
			}).Info("Not found")

			ctx.StatusCode(iris.StatusNotFound)
			ctx.JSON(iris.Map{
				"status":  "failed",
				"message": "Not found",
				"payload": &controller.LastResult{},
			})
			return
		}

		log.WithFields(logrus.Fields{
			"repository": URLparams["repository"],
			"branch":     URLparams["branch"],
		}).Info("Reports will be returned")

		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(iris.Map{
			"status":  "success",
			"message": "Found",
			"payload": reports,
		})
	})

	app.Get("/reports/total", func(ctx iris.Context) {
		URLparams := ctx.URLParams()

		params := controller.GetArgs{
			Repository: URLparams["repository"],
			Branch:     URLparams["branch"],
		}

		reports, err := ctrl.Total(&params)

		if err != nil {
			setGetError(log, &params, ctx, err)
			return
		}

		if len(reports) == 0 {
			log.WithFields(logrus.Fields{
				"repository": URLparams["repository"],
				"branch":     URLparams["branch"],
			}).Info("Not found")

			ctx.StatusCode(iris.StatusNotFound)
			ctx.JSON(iris.Map{
				"status":  "failed",
				"message": "Not found",
				"payload": &controller.LastResult{},
			})
			return
		}

		log.WithFields(logrus.Fields{
			"repository": URLparams["repository"],
			"branch":     URLparams["branch"],
		}).Info("Total report will be returned")

		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(iris.Map{
			"status":  "success",
			"message": "Found",
			"payload": reports,
		})
	})
}
