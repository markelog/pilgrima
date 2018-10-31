package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kataras/iris"
	"github.com/markelog/pilgrima/application"
	"github.com/markelog/pilgrima/database"
	"github.com/markelog/pilgrima/log"
	"github.com/sirupsen/logrus"
)

func main() {
	log := log.Log()

	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
	}

	var (
		port    = os.Getenv("PORT")
		address = ":" + port

		db  = database.Up()
		app = application.Up(db)
	)

	log.WithFields(logrus.Fields{
		"address": address,
	}).Info("Started")
	app.Run(iris.Addr(address))
}
