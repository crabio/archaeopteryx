package example

import (
	// External

	"github.com/jinzhu/configor"
	"github.com/sirupsen/logrus"

	// Internal
	"github.com/iakrevetkho/archaeopteryx"
	api_hello_world_v1 "github.com/iakrevetkho/archaeopteryx/example/pkg/api/hello_world/v1"
	api_data "github.com/iakrevetkho/archaeopteryx/pkg/api/data"
	archaeopteryx_config "github.com/iakrevetkho/archaeopteryx/pkg/config"
)

var (
	externalGrpcServicesRegistrars = []archaeopteryx.ExternalGrpcServiceRegistrar{
		api_hello_world_v1.RegisterServiceServer,
	}
	externalGrpcProxyServicesRegistrars = []archaeopteryx.ExternalGrpcProxyServiceRegistrar{
		api_hello_world_v1.RegisterProxyServiceServer,
	}
)

func main() {
	log := logrus.WithField("component", "main")

	// Init archeopteryx config
	conf := new(archaeopteryx_config.Config)
	if err := configor.Load(conf, "config.yml"); err != nil {
		log.WithError(err).Fatal("couldn't init config")
	}

	// Init your controller
	controllers := api_data.Controllers{}

	// Create archeopteryx server
	server, err := archaeopteryx.New(conf, externalGrpcServicesRegistrars, externalGrpcProxyServicesRegistrars, controllers)
	if err != nil {
		log.WithError(err).Fatal("couldn't init server")
	}

	// Run archeopteryx server
	if err := server.Run(); err != nil {
		log.WithError(err).Fatal("couldn't run server")
	}
}
