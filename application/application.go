package application

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recover"
	"github.com/markelog/pilgrima/application/root"
	"github.com/markelog/pilgrima/application/token"
	"github.com/markelog/pilgrima/log"
)

// Up app
func Up(db *gorm.DB) *iris.Application {
	var (
		application = iris.New()
		log         = log.Log()
	)

	application.Logger().Install(log)
	application.Use(recover.New())

	root.Set(application, db)
	token.Set(application, db)

	application.Configure(iris.WithConfiguration(iris.Configuration{
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

	return application
}
