package archaeopteryx

import (
	// External
	"net"
	"strconv"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	// Internal
	api_data "github.com/iakrevetkho/archaeopteryx/pkg/api/data"
	"github.com/iakrevetkho/archaeopteryx/pkg/helpers"
)

type grpcServer struct {
	log        *logrus.Entry
	port       int
	grpcServer *grpc.Server
}

// Function creates gRPC server on the [port]
func newGrpcServer(port int, controllers *api_data.Controllers, services []IServiceServer) (*grpcServer, error) {
	s := new(grpcServer)
	s.log = helpers.CreateComponentLogger("archeaopteryx-grpc")
	s.port = port
	s.grpcServer = grpc.NewServer()

	// Register service routes
	for _, service := range services {
		if err := service.RegisterGrpc(s.grpcServer); err != nil {
			return nil, err
		}
	}
	s.log.Debug("Services are registered")

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
