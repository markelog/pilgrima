package main

import (
	"os"

	"github.com/kataras/iris/v12"
	"github.com/markelog/pilgrima/database"
	"github.com/markelog/pilgrima/env"
	"github.com/markelog/pilgrima/logger"
	"github.com/markelog/pilgrima/routes"
	"github.com/markelog/pilgrima/routes/common"
	"github.com/markelog/pilgrima/routes/projects"
	"github.com/markelog/pilgrima/routes/reports"
	"github.com/markelog/pilgrima/routes/root"
	"github.com/markelog/pilgrima/routes/tokens"
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

	defer db.Close()

	root.Up(app, db, log)
	tokens.Up(app, db, log)
	projects.Up(app, db, log)
	reports.Up(app, db, log)
	common.Up(app, db, log)

	log.WithFields(logrus.Fields{
		"port": port,
	}).Info("Started")
	app.Run(iris.Addr(address))
}
