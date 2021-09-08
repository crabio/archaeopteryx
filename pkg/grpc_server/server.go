package grpc_server

import (
	// External

	"net"
	"strconv"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"

	// Internal
	"github.com/iakrevetkho/archaeopteryx/config"
	api_data "github.com/iakrevetkho/archaeopteryx/pkg/api/data"
	"github.com/iakrevetkho/archaeopteryx/pkg/helpers"
	"github.com/iakrevetkho/archaeopteryx/service"
)

type Server struct {
	log        *logrus.Entry
	addr       string
	grpcServer *grpc.Server
	listener   net.Listener
}

func New(c *config.Config, controllers *api_data.Controllers, services []service.IServiceServer) (*Server, error) {
	s := new(Server)
	s.log = helpers.CreateComponentLogger("archeaopteryx-grpc")
	s.addr = ":" + strconv.FormatUint(c.GrpcPort, 10)

	// Check that we have FS with certificates
	if c.Secutiry.TlsConfig != nil {
		s.log.Info("Create gRPC server with TLS security")
		s.grpcServer = grpc.NewServer(grpc.Creds(credentials.NewTLS(c.Secutiry.TlsConfig)))
	} else {
		s.log.Info("Create insecure gRPC server")
		s.grpcServer = grpc.NewServer()
	}

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

func (s *Server) Run() error {
	var err error

	s.listener, err = net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	go func() {
		if err := s.grpcServer.Serve(s.listener); err != nil {
			s.log.WithError(err).Fatal("Couldn't serve gRPC server")
		}
	}()
	s.log.WithField("url", s.addr).Info("Serving gRPC")

	return nil
}

func (s *Server) Stop() {
	s.grpcServer.Stop()
}
