package api_health_v1

import (
	// External

	"time"

	"github.com/alexliesenfeld/health"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	// Internal
	api_data "github.com/iakrevetkho/archaeopteryx/pkg/api/data"
	"github.com/iakrevetkho/archaeopteryx/pkg/helpers"
	health_v1 "github.com/iakrevetkho/archaeopteryx/proto/gen/health/v1"
)

const (
	watchUpdatePeriod = time.Second * 15
)

type HealthServiceServer struct {
	log *logrus.Entry
	// Required to have revese compatability
	health_v1.UnimplementedHealthServer
	checker health.Checker
}

func RegisterServiceServer(s grpc.ServiceRegistrar, controllers *api_data.Controllers) error {
	server := new(HealthServiceServer)
	server.log = helpers.CreateComponentLogger("grpc-healthcheck-v1")
	server.checker = controllers.HealthChecker

	// Attach the Health service to the server
	health_v1.RegisterHealthServer(s, server)
	server.log.Debug("Service registered")

	return nil
}
