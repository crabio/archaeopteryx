package archaeopteryx

import (
	// External
	"net"
	"strconv"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	// Internal
	api_data "github.com/iakrevetkho/archaeopteryx/pkg/api/data"
	api_health_v1 "github.com/iakrevetkho/archaeopteryx/pkg/api/health/v1"
	"github.com/iakrevetkho/archaeopteryx/pkg/helpers"
)

type internalGrpcServiceRegistrar func(registrar grpc.ServiceRegistrar, controllers *api_data.Controllers) error
type ExternalGrpcServiceRegistrar func(externalRegistrar grpc.ServiceRegistrar, externalControllers interface{}) error

var (
	internalGrpcServicesRegistrars = []internalGrpcServiceRegistrar{
		api_health_v1.RegisterServiceServer,
	}
)

type grpcServer struct {
	log        *logrus.Entry
	port       int
	grpcServer *grpc.Server
}

// Function creates gRPC server on the [port]
func newGrpcServer(port int, controllers *api_data.Controllers, externalServicesRegistrars []ExternalGrpcServiceRegistrar, externalControllers interface{}) (*grpcServer, error) {
	s := new(grpcServer)
	s.log = helpers.CreateComponentLogger("grpc")
	s.port = port
	s.grpcServer = grpc.NewServer()

	// Register internal service routes
	for _, serviceRegistrars := range internalGrpcServicesRegistrars {
		if err := serviceRegistrars(s.grpcServer, controllers); err != nil {
			return nil, err
		}
	}
	s.log.Debug("Internal services are registered")

	// Register external service routes
	for _, serviceRegistrars := range externalServicesRegistrars {
		if err := serviceRegistrars(s.grpcServer, externalControllers); err != nil {
			return nil, err
		}
	}
	s.log.Debug("External services are registered")

	return s, nil
}

// Function runs gRPC server on the [port]
func (s *grpcServer) run() error {
	// Create a listener on TCP port
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(s.port))
	if err != nil {
		return err
	}

	go func() {
		if err := s.grpcServer.Serve(listener); err != nil {
			s.log.WithError(err).Fatal("Couldn't serve gRPC server")
		}
	}()
	s.log.WithField("url", ":"+strconv.Itoa(s.port)).Debug("Serving gRPC")

	return nil
}
