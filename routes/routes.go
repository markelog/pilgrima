package routes

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/middleware/recover"
	"github.com/markelog/pilgrima/logger"
)

// Up app
func Up() *iris.Application {
	var (
		app = iris.New()
		log = logger.Up()
	)

	app.Logger().Install(log)
	app.Use(crs())
	app.Use(recover.New())

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

	return app
}

func crs() context.Handler {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
	})
}
