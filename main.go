package main

import (
	// External
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"

	// Internal
	"github.com/iakrevetkho/archaeopteryx/pkg/api"
	"github.com/iakrevetkho/archaeopteryx/pkg/helpers"
)

func main() {
	helpers.InitLogger()
	logrus.SetLevel(logrus.TraceLevel)

	log := helpers.CreateComponentLogger("main")

	api.RunServer()

	log.Info("Wait exit signal")
	quitSignal := make(chan os.Signal, 1)
	signal.Notify(quitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-quitSignal
}
