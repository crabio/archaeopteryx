package grpc_server

import (
	// External

	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// Internal
	"github.com/iakrevetkho/archaeopteryx/pkg/helpers"
	"github.com/iakrevetkho/archaeopteryx/service"
)

type Server struct {
	log        *logrus.Entry
	grpcServer *grpc.Server
}

func New(services []service.IServiceServer) (*Server, error) {
	s := new(Server)
	s.log = helpers.CreateComponentLogger("archeaopteryx-grpc")
	s.grpcServer = grpc.NewServer()

	// Register service routes
	for _, service := range services {
		if err := service.RegisterGrpc(s.grpcServer); err != nil {
			return nil, err
		}
	}
	s.log.Debug("Services are registered")

	reflection.Register(s.grpcServer)
	s.log.Debug("Reflection service on gRPC server is registered")

	return s, nil
}

func (s *Server) Run(l net.Listener) error {
	go func() {
		if err := s.grpcServer.Serve(l); err != nil {
			s.log.WithError(err).Fatal("Couldn't serve gRPC server")
		}
	}()
	s.log.Info("Serving gRPC")

	return nil
}
