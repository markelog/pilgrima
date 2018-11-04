package token

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	controller "github.com/markelog/pilgrima/controllers/token"
)

type postProject struct {
	Project int `json:"project"`
}

// Up token route
func Up(app *iris.Application, db *gorm.DB) {
	app.Post("/token", func(ctx iris.Context) {
		var params postProject
		ctx.ReadJSON(&params)

		token, err := controller.New(params.Project, db).Create()

		if err == controller.ErrNoSuchProject {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{
				"status":  "failed",
				"message": "Can't find this project",
				"payload": iris.Map{},
			})
			return
		}

		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(iris.Map{
			"status":  "created",
			"message": "Yey!",
			"payload": iris.Map{
				"token": token,
			},
		})
	})
}
