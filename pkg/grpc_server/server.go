package grpc_server

import (
	// External

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// Internal
	api_data "github.com/iakrevetkho/archaeopteryx/pkg/api/data"
	"github.com/iakrevetkho/archaeopteryx/pkg/helpers"
	"github.com/iakrevetkho/archaeopteryx/service"
)

type Server struct {
	log        *logrus.Entry
	grpcServer *grpc.Server
}

func New(controllers *api_data.Controllers, services []service.IServiceServer) (*Server, error) {
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

func (s *Server) GetHttpHandler() gin.HandlerFunc {
	return gin.WrapH(s.grpcServer)
}
