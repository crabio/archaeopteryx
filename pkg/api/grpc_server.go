package api

import (
	// External
	"net"
	"strconv"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	// Internal
	api_data "github.com/iakrevetkho/archaeopteryx/pkg/api/data"
	api_healthcheck_v1 "github.com/iakrevetkho/archaeopteryx/pkg/api/healthcheck/v1"
	api_hello_world_v1 "github.com/iakrevetkho/archaeopteryx/pkg/api/hello_world/v1"
	api_user_v1 "github.com/iakrevetkho/archaeopteryx/pkg/api/user/v1"
	api_user_v2 "github.com/iakrevetkho/archaeopteryx/pkg/api/user/v2"
	"github.com/iakrevetkho/archaeopteryx/pkg/helpers"
)

var (
	// List with all service registars.
	// If you add new service, you need to add registrar here.
	grpcServicesRegistrars = []func(grpc.ServiceRegistrar, *api_data.Controllers) error{
		api_healthcheck_v1.RegisterServiceServer,
		api_hello_world_v1.RegisterServiceServer,
		api_user_v1.RegisterServiceServer,
		api_user_v2.RegisterServiceServer,
	}
)

type grpcServer struct {
	log        *logrus.Entry
	port       int
	grpcServer *grpc.Server
}

// Function creates gRPC server on the [port]
func newGrpcServer(port int, controllers *api_data.Controllers) (*grpcServer, error) {
	server := new(grpcServer)
	server.log = helpers.CreateComponentLogger("grpc")
	server.port = port
	server.grpcServer = grpc.NewServer()

	// Register service routes
	for _, serviceRestrar := range grpcServicesRegistrars {
		if err := serviceRestrar(server.grpcServer, controllers); err != nil {
			return nil, err
		}
	}
	server.log.Debug("Services are registered")

	return server, nil
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
