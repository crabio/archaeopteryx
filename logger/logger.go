package logger

import (
	// External
	"github.com/sirupsen/logrus"
	// Internal
	"github.com/iakrevetkho/archaeopteryx/pkg/helpers"
)

// CreateLogger - creates logger bases from archaeopteryx base logger,
// which has formatter and writes data to thi ratotated log file.
func CreateLogger(componentName string) *logrus.Entry {
	logger := helpers.CreateComponentLogger(componentName)
	return logger
}
