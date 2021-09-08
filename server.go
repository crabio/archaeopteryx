package archaeopteryx

import (
	// External

	"fmt"
	"os"
	"os/signal"
	"syscall"

	// Internal
	"github.com/iakrevetkho/archaeopteryx/config"
	api_data "github.com/iakrevetkho/archaeopteryx/pkg/api/data"
	api_health_v1 "github.com/iakrevetkho/archaeopteryx/pkg/api/health/v1"
	"github.com/iakrevetkho/archaeopteryx/pkg/grpc_proxy_server"
	"github.com/iakrevetkho/archaeopteryx/pkg/grpc_server"
	"github.com/iakrevetkho/archaeopteryx/pkg/healthchecker"
	"github.com/iakrevetkho/archaeopteryx/pkg/helpers"
	"github.com/iakrevetkho/archaeopteryx/pkg/http"
	"github.com/iakrevetkho/archaeopteryx/pkg/swagger"
	"github.com/iakrevetkho/archaeopteryx/service"
	"github.com/sirupsen/logrus"
)

type Server struct {
	Config *config.Config

	log         *logrus.Entry
	controllers *api_data.Controllers
	services    []service.IServiceServer

	grpcs      *grpc_server.Server
	grpcps     *grpc_proxy_server.Server
	httpServer *http.Server
}

func New(config *config.Config, externalServices []service.IServiceServer) (*Server, error) {
	var err error

	s := new(Server)

	helpers.InitLogger(config)
	s.log = helpers.CreateComponentLogger("archeaopteryx-server")
	s.log.WithField("config", helpers.MustMarshal(config)).Info("Config is inited")

	s.Config = config

	// Read TLS config
	if s.Config.Secutiry.CertFS != nil &&
		s.Config.Secutiry.CertName != nil &&
		s.Config.Secutiry.KeyName != nil {
		s.Config.Secutiry.TlsConfig, err = helpers.CreateTlsConfig(s.Config)
		if err != nil {
			return nil, fmt.Errorf("couldn't create TLS config. " + err.Error())
		}
	} else {
		s.log.Warn("You haven't specified TLS certificates")
	}

	s.controllers = new(api_data.Controllers)
	s.controllers.Config = config
	s.controllers.HealthChecker = healthchecker.New()

	// Add internal services
	s.services = append(s.services, api_health_v1.New(s.controllers))

	// Add external services
	s.services = append(s.services, externalServices...)

	s.grpcs, err = grpc_server.New(s.Config, s.controllers, s.services)
	if err != nil {
		return nil, fmt.Errorf("couldn't create gRPC server. " + err.Error())
	}

	// Init gRPC server proxy on run, because it can be inited only with working gRPC server
	s.grpcps = grpc_proxy_server.New(s.Config)
	if err != nil {
		return nil, fmt.Errorf("couldn't create gRPC proxy server. " + err.Error())
	}

	// Run gRPC server before creating gRPC proxy to allow gRPC proxy dial connection with gRPC
	if err := s.grpcs.Run(); err != nil {
		return nil, fmt.Errorf("couldn't run gRPC server. " + err.Error())
	}

	if err := s.grpcps.RegisterServices(s.services); err != nil {
		return nil, fmt.Errorf("couldn't register gRPC proxy services. " + err.Error())
	}

	sws, err := swagger.New(s.Config)
	if err != nil {
		return nil, fmt.Errorf("couldn't create Swagger docs server. " + err.Error())
	}

	s.httpServer = http.New(s.Config, s.grpcps, sws)

	return s, nil
}

func (s *Server) Run() error {
	s.httpServer.Run()

	s.log.Info("Wait exit signal")
	quitSignal := make(chan os.Signal, 1)
	signal.Notify(quitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-quitSignal

	return nil
}
