package helpers

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

func CreateComponentLogger(componentName string) *logrus.Entry {
	log := logrus.New()
	log.SetFormatter(&nested.Formatter{})
	return log.WithField("component", componentName)
}
