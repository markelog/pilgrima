package main

import (
	"os"

	// Routes
	"github.com/markelog/pilgrima/routes/common"
	"github.com/markelog/pilgrima/routes/project"
	"github.com/markelog/pilgrima/routes/report"
	"github.com/markelog/pilgrima/routes/root"
	"github.com/markelog/pilgrima/routes/token"

	// Internal dependencies
	"github.com/markelog/pilgrima/database"
	"github.com/markelog/pilgrima/env"
	"github.com/markelog/pilgrima/logger"
	"github.com/markelog/pilgrima/routes"

	// External dependencies
	"github.com/kataras/iris"
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
	project.Up(app, db, log)
	report.Up(app, db, log)
	common.Up(app, db, log)

	log.WithFields(logrus.Fields{
		"port": port,
	}).Info("Started")
	app.Run(iris.Addr(address))
}
