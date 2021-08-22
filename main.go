package main

import (
	// External

	"os"
	"os/signal"
	"syscall"

	// Internal
	"github.com/iakrevetkho/archaeopteryx/pkg/api"
	"github.com/iakrevetkho/archaeopteryx/pkg/config"
	"github.com/iakrevetkho/archaeopteryx/pkg/helpers"
)

func main() {
	conf := config.LoadConfig()
	helpers.InitLogger(conf)
	log := helpers.CreateComponentLogger("main")
	log.WithField("config", helpers.MustMarshal(conf)).Info("Config is inited")

	api.RunServer()

	log.Info("Wait exit signal")
	quitSignal := make(chan os.Signal, 1)
	signal.Notify(quitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-quitSignal
}
