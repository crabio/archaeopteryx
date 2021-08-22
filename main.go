package main

import (
	// External
	"os"
	"os/signal"
	"syscall"

	// Internal
	"github.com/iakrevetkho/archaeopteryx/pkg/api"
	"github.com/iakrevetkho/archaeopteryx/pkg/helpers"
)

func main() {
	log := helpers.CreateComponentLogger("main")

	api.RunServer()

	log.Info("Wait exit signal")
	quitSignal := make(chan os.Signal, 1)
	signal.Notify(quitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-quitSignal
}
