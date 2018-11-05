package main

import (
	"os"

	"github.com/kataras/iris"
	"github.com/markelog/pilgrima/database"
	"github.com/markelog/pilgrima/env"
	"github.com/markelog/pilgrima/logger"
	"github.com/markelog/pilgrima/routes"
	"github.com/markelog/pilgrima/routes/root"
	"github.com/markelog/pilgrima/routes/token"
	"github.com/sirupsen/logrus"
)

func main() {
	env.Up()

	var (
		port    = os.Getenv("PORT")
		address = ":" + port
	)

	var (
		app = routes.Up()
		db  = database.Up()
		log = logger.Up()
	)

	root.Up(app, db, log)
	token.Up(app, db, log)

	log.WithFields(logrus.Fields{
		"port": port,
	}).Info("Started")
	app.Run(iris.Addr(address))
}
