package main

import (
	"os"

	"github.com/kataras/iris"
	"github.com/markelog/pilgrima/application"
	"github.com/markelog/pilgrima/application/root"
	"github.com/markelog/pilgrima/application/token"
	"github.com/markelog/pilgrima/database"
	"github.com/markelog/pilgrima/env"
	"github.com/markelog/pilgrima/logger"
	"github.com/sirupsen/logrus"
)

func main() {
	env.Up()

	var (
		port    = os.Getenv("PORT")
		address = ":" + port
	)

	var (
		app = application.Up()
		db  = database.Up()
		log = logger.Up()
	)

	root.Up(app, db)
	token.Up(app, db)

	log.WithFields(logrus.Fields{
		"port": port,
	}).Info("Started")
	app.Run(iris.Addr(address))
}
