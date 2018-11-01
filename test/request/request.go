package request

import (
	"net/http"
	"testing"

	"github.com/kataras/iris"
	httpexpect "gopkg.in/gavv/httpexpect.v1"
)

// Up request
func Up(app *iris.Application, t *testing.T) *httpexpect.Expect {
	return httpexpect.WithConfig(httpexpect.Config{
		BaseURL: "http://example.com",
		Client: &http.Client{
			Transport: httpexpect.NewBinder(app),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewRequireReporter(t),
	})
}
