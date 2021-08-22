package helpers

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

func InitLogger(level logrus.Level) {
	logrus.SetLevel(level)
	logrus.SetFormatter(&nested.Formatter{})
}

func CreateComponentLogger(componentName string) *logrus.Entry {
	return logrus.WithField("component", componentName)
}
