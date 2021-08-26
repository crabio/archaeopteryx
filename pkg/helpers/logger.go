package helpers

import (
	// External
	"io"
	"os"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/iakrevetkho/archaeopteryx/pkg/config"
	"github.com/iakrevetkho/woodpecker"
	"github.com/sirupsen/logrus"
)

const LOG_PREFIX = "archeaopteryx"

func InitLogger(conf *config.Config) {
	logrus.SetLevel(conf.Log.Level)
	logrus.SetFormatter(&nested.Formatter{})

	rotatedLogFile := woodpecker.New(woodpecker.Config{
		Filename:       conf.Log.Filename,
		RotateEveryday: true,
		MaxAgeInDays:   conf.Log.MaxAgeInDays,
		MaxSizeInMb:    conf.Log.MaxSizeInMb,
		MaxBackups:     conf.Log.MaxBackups,
		Compress:       conf.Log.Compress,
	})

	mw := io.MultiWriter(os.Stdout, rotatedLogFile)

	logrus.SetOutput(mw)
}

func CreateComponentLogger(componentName string) *logrus.Entry {
	return logrus.WithField("component", LOG_PREFIX+"-"+componentName)
}
