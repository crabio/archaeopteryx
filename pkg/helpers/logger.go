package helpers

import (
	// External
	"io"
	"os"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/iakrevetkho/archaeopteryx/pkg/config"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger(conf *config.Config) {
	logrus.SetLevel(conf.Log.Level)
	logrus.SetFormatter(&nested.Formatter{})

	rotatedLogFile := &lumberjack.Logger{
		Filename:   conf.Log.Filename,
		MaxAge:     conf.Log.MaxAgeInDays,
		MaxSize:    conf.Log.MaxSizeInMb,
		MaxBackups: conf.Log.MaxBackups,
		Compress:   conf.Log.Compress,
	}

	mw := io.MultiWriter(os.Stdout, rotatedLogFile)

	logrus.SetOutput(mw)

}

func CreateComponentLogger(componentName string) *logrus.Entry {
	return logrus.WithField("component", componentName)
}
