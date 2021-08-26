package main

import (
	// External

	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"

	// Internal
	"github.com/iakrevetkho/archaeopteryx/pkg/api"
	api_data "github.com/iakrevetkho/archaeopteryx/pkg/api/data"
	"github.com/iakrevetkho/archaeopteryx/pkg/config"
	"github.com/iakrevetkho/archaeopteryx/pkg/healthchecker"
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

	// Init controllers
	controllers := new(api_data.Controllers)
	controllers.HealthChecker = healthchecker.New()

	// Run API server
	api.RunServer(controllers)

	log.Info("Wait exit signal")
	quitSignal := make(chan os.Signal, 1)
	signal.Notify(quitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-quitSignal
}
