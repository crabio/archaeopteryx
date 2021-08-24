package api_healthcheck_v1

import (
	// External
	"time"

	"github.com/alexliesenfeld/health"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	// Internal
	api_data "github.com/iakrevetkho/archaeopteryx/pkg/api/data"
	"github.com/iakrevetkho/archaeopteryx/pkg/helpers"
	healthcheck_v1 "github.com/iakrevetkho/archaeopteryx/proto/gen/healthcheck/v1"
)

const (
	watchUpdatePeriod = time.Second * 15
)

type HealthcheckServiceServer struct {
	log *logrus.Entry
	// Required to have revese compatability
	healthcheck_v1.UnimplementedHealthCheckServiceServer
	checker health.Checker
}

func RegisterServiceServer(s grpc.ServiceRegistrar, controllers *api_data.Controllers) error {
	server := new(HealthcheckServiceServer)
	server.log = helpers.CreateComponentLogger("grpc-healthcheck-v1")
	server.checker = controllers.HealthChecker

	// Attach the Healthcheck service to the server
	healthcheck_v1.RegisterHealthCheckServiceServer(s, server)
	server.log.Debug("Service registered")

	return nil
}
