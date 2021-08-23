package main

import (
	// External

	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"

	// Internal
	"github.com/iakrevetkho/archaeopteryx/pkg/api"
	"github.com/iakrevetkho/archaeopteryx/pkg/config"
	"github.com/iakrevetkho/archaeopteryx/pkg/helpers"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		logrus.WithError(err).Fatal("couldn't init config")
	}

	helpers.InitLogger(conf)
	log := helpers.CreateComponentLogger("main")
	log.WithField("config", helpers.MustMarshal(conf)).Info("Config is inited")

	api.RunServer()

	log.Info("Wait exit signal")
	quitSignal := make(chan os.Signal, 1)
	signal.Notify(quitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-quitSignal
}
