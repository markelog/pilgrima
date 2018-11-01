package token

import (
	"crypto/rand"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/markelog/pilgrima/database/models"
)

type postProject struct {
	Project int `json:"project"`
}

var token models.Token
var project models.Project

func generate() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

// Up token route
func Up(app *iris.Application, db *gorm.DB) {
	app.Post("/token", func(ctx iris.Context) {
		var params postProject
		ctx.ReadJSON(&params)

		db.Where("id = ?", params.Project).First(&project)

		if project.ID == 0 {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{
				"status":  "failed",
				"message": "Can't find this project",
				"payload": iris.Map{},
			})
			return
		}

		var token = &models.Token{
			Token:   generate(),
			Project: project,
		}

		db.Create(&token)

		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(iris.Map{
			"status":  "created",
			"message": "Yey!",
			"payload": iris.Map{
				"token": token.Token,
			},
		})
	})
}
