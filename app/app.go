package app

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recover"
	"github.com/markelog/pilgrima/log"
	"github.com/sirupsen/logrus"
)

func Start(address string, db *gorm.DB) {
	log := log.Log()
	app := iris.New()
	app.Logger().Install(log)
	app.Use(recover.New())

	Root(app, db)

	app.Configure(iris.WithConfiguration(iris.Configuration{
		DisableStartupLog:                 true,
		DisableInterruptHandler:           false,
		DisablePathCorrection:             false,
		EnablePathEscape:                  false,
		FireMethodNotAllowed:              false,
		DisableBodyConsumptionOnUnmarshal: false,
		DisableAutoFireStatusCode:         false,
		TimeFormat:                        "Mon, 02 Jan 2006 15:04:05 GMT",
		Charset:                           "UTF-8",
	}))

	log.WithFields(logrus.Fields{
		"address": address,
	}).Info("Started")

	app.Run(iris.Addr(address))
}
