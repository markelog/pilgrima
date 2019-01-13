package projects

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	controller "github.com/markelog/pilgrima/controllers/project"
	"github.com/sirupsen/logrus"
	"github.com/xeipuuv/gojsonschema"
)

type postProject struct {
	Name       string `json:"name"`
	Repository string `json:"repository"`
}

var schema = gojsonschema.NewStringLoader(`{
	"type": "object",
	"properties": {
		"name": {"type": "string", "minLength": 1},
		"repository": {"type": "string", "minLength": 1}
	},
	"required": ["name", "repository"]
}`)

func validate(params *postProject) (*gojsonschema.Result, *iris.Map) {
	var (
		paramsLoader = gojsonschema.NewGoLoader(params)
		check, _     = gojsonschema.Validate(schema, paramsLoader)

		errors  []string
		payload *iris.Map
	)

	if check.Valid() == false {
		for _, desc := range check.Errors() {
			errors = append(errors, desc.String())
		}

		payload = &iris.Map{"errors": errors}

		return check, payload
	}

	return check, nil
}

// Up project route
func Up(app *iris.Application, db *gorm.DB, log *logrus.Logger) {
	app.Post("/projects", func(ctx iris.Context) {
		var params postProject
		ctx.ReadJSON(&params)

		validation, errors := validate(&params)

		if validation.Valid() == false {
			log.WithFields(logrus.Fields{
				"project":    params.Name,
				"repository": params.Repository,
			}).Error("Params are not valid")

			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{
				"status":  "failed",
				"message": "Params are not valid",
				"payload": errors,
			})

			return
		}

		ctrl := controller.New(db)
		result, err := ctrl.Create(params.Name, params.Repository)

		if err != nil {
			log.WithFields(logrus.Fields{
				"project":    params.Name,
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

		log.WithFields(logrus.Fields{
			"project": params.Name,
		}).Info("Project created")

		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(iris.Map{
			"status":  "created",
			"message": "Yey!",
			"payload": iris.Map{
				"project": result.Name,
				"id":      result.ID,
			},
		})
	})

	app.Get("/projects", func(ctx iris.Context) {
		ctrl := controller.New(db)
		projects, err := ctrl.List()

		if err != nil {
			errorString := err.Error()
			log.Error(errorString)

			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{
				"status":  "failed",
				"message": "Something went wrong",
				"payload": iris.Map{},
			})

			return
		}

		if len(projects) == 0 {
			log.Error("Can't find any projects")

			ctx.StatusCode(iris.StatusNotFound)
			ctx.JSON(iris.Map{
				"status":  "failed",
				"message": "Not found",
				"payload": iris.Map{},
			})

			return
		}

		log.Info("Projects returned")

		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(iris.Map{
			"status":  "success",
			"message": "There you go",
			"payload": projects,
		})

		return
	})
}
