package main

import (
	// External
	"github.com/jinzhu/configor"
	"github.com/sirupsen/logrus"

	// Internal
	"github.com/iakrevetkho/archaeopteryx"
	archaeopteryx_config "github.com/iakrevetkho/archaeopteryx/config"
)

func main() {
	log := logrus.WithField("component", "main")

	// Init archeopteryx config
	conf := new(archaeopteryx_config.Config)
	if err := configor.Load(conf, "config.yml"); err != nil {
		log.WithError(err).Fatal("couldn't init config")
	}

	// Init services
	services := []archaeopteryx.IServiceServer{}

	// Create archeopteryx server
	server := archaeopteryx.New(conf, services)

	// Run archeopteryx server
	log.Info("Run server")
	if err := server.Run(); err != nil {
		log.WithError(err).Fatal("couldn't run server")
	}
}
