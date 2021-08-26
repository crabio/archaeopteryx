package archaeopteryx

import (
	// External
	"fmt"
	"os"
	"os/signal"
	"syscall"

	// Internal
	api_data "github.com/iakrevetkho/archaeopteryx/pkg/api/data"
	"github.com/iakrevetkho/archaeopteryx/pkg/config"
	"github.com/iakrevetkho/archaeopteryx/pkg/healthchecker"
	"github.com/iakrevetkho/archaeopteryx/pkg/helpers"
	"github.com/sirupsen/logrus"
)

type Server struct {
	Config *config.Config

	log             *logrus.Entry
	controllers     *api_data.Controllers
	grpcServer      *grpcServer
	grpcProxyServer *grpcProxyServer
}

func New(config *config.Config, externalGrpcServicesRegistrars []ExternalGrpcServiceRegistrar, externalGrpcProxyServicesRegistrars []ExternalGrpcProxyServiceRegistrar, externalControllers interface{}) (*Server, error) {
	var err error

	s := new(Server)

	helpers.InitLogger(config)
	s.log = helpers.CreateComponentLogger("server")
	s.log.WithField("config", helpers.MustMarshal(config)).Info("Config is inited")

	s.Config = config
	s.controllers = new(api_data.Controllers)
	s.controllers.HealthChecker = healthchecker.New()
	s.grpcServer, err = newGrpcServer(s.Config.GrpcPort, s.controllers, externalGrpcServicesRegistrars, externalControllers)
	if err != nil {
		return nil, fmt.Errorf("couldn't create gRPC server. " + err.Error())
	}
	s.grpcProxyServer, err = newGrpcProxyServer(s.Config.GrpcGatewayPort, s.grpcServer, s.controllers, externalGrpcProxyServicesRegistrars, externalControllers)
	if err != nil {
		return nil, fmt.Errorf("couldn't create gRPC proxy server. " + err.Error())
	}

	return s, nil
}

func (s *Server) Run() error {
	if err := s.grpcServer.run(); err != nil {
		return fmt.Errorf("couldn't run gRPC server. " + err.Error())
	}
	if err := s.grpcProxyServer.run(); err != nil {
		return fmt.Errorf("couldn't run gRPC proxy server. " + err.Error())
	}

	s.log.Info("Wait exit signal")
	quitSignal := make(chan os.Signal, 1)
	signal.Notify(quitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-quitSignal

	return nil
}
