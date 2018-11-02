package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Up logger
func Up() *logrus.Logger {
	var log = logrus.New()

	log.Out = os.Stdout

	return log
}
