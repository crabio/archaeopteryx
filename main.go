package main

import (
	// External
	"log"
	"os"
	"os/signal"
	"syscall"

	// Internal
	"github.com/iakrevetkho/archaeopteryx/pkg/api"
)

func main() {
	api.NewServer()

	log.Println("Wait exit signal")
	quitSignal := make(chan os.Signal, 1)
	signal.Notify(quitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-quitSignal
}
