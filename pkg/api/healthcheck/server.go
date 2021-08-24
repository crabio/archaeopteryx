package api_healthcheck_v1

import (
	// External
	"github.com/alexliesenfeld/health"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	// Internal
	"github.com/iakrevetkho/archaeopteryx/pkg/helpers"
	healthcheck_v1 "github.com/iakrevetkho/archaeopteryx/proto/healthcheck/v1"
)

type HealthcheckServiceServer struct {
	log *logrus.Entry
	// Required to have revese compatability
	healthcheck_v1.UnimplementedHealthCheckServer
	checker health.Checker
}

func RegisterServiceServer(s grpc.ServiceRegistrar, checker health.Checker) error {
	server := new(HealthcheckServiceServer)
	server.log = helpers.CreateComponentLogger("grpc-healthcheck-v1")
	server.checker = checker

	// Attach the Healthcheck service to the server
	healthcheck_v1.RegisterHealthCheckServer(s, &HealthcheckServiceServer{})
	server.log.Debug("Service registered")

	return nil
}
