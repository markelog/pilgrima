package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

func Log() *logrus.Logger {
	var log = logrus.New()

	log.Out = os.Stdout

	return log
}
