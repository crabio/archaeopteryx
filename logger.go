package archaeopteryx

import (
	// External
	"github.com/sirupsen/logrus"
	// Internal
	"github.com/iakrevetkho/archaeopteryx/pkg/helpers"
)

type Logger logrus.Entry

// CreateLogger - creates logger bases from archaeopteryx base logger,
// which has formatter and writes data to thi ratotated log file.
func CreateLogger(componentName string) *Logger {
	logger := Logger(*helpers.CreateComponentLogger(componentName))
	return &logger
}
