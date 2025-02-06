package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func NewLogger() {
	if Logger != nil {
		return
	}

	Logger = logrus.New()
	Logger.SetOutput(os.Stdout)
}
