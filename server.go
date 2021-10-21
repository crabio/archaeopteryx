package archaeopteryx

import (
	// External
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/soheilhy/cmux"

	// Internal
	"github.com/iakrevetkho/archaeopteryx/config"
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
	log      *logrus.Entry
	services []service.IServiceServer

	// Main TCP listener
	listener net.Listener
	// gRPC sub-listener
	grpcListener net.Listener
	// HTTP sub-listener
	httpListener net.Listener

	grpcs      *grpc_server.Server
	grpcps     *grpc_proxy_server.Server
	httpServer *http.Server
}

func New(cfg *config.Config, externalServices []service.IServiceServer) (*Server, error) {
	var err error

	s := new(Server)

	helpers.InitLogger(cfg)
	s.log = helpers.CreateComponentLogger("archeaopteryx-server")
	s.log.WithField("config", helpers.MustMarshal(cfg)).Info("Config is inited")

	if s.listener, err = net.Listen("tcp", ":8080"); err != nil {
		return nil, fmt.Errorf("couldn't create net listener. " + err.Error())
	}
	m := cmux.New(s.listener)
	s.grpcListener = m.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	s.httpListener = m.Match(cmux.HTTP1Fast())

	// Add internal services
	s.services = append(s.services, api_health_v1.New(healthchecker.New(), cfg.Health.WatchUpdatePeriod))

	// Add external services
	s.services = append(s.services, externalServices...)

	s.grpcs, err = grpc_server.New(s.services)
	if err != nil {
		return nil, fmt.Errorf("couldn't create gRPC server. " + err.Error())
	}

	// Init gRPC server proxy on run, because it can be inited only with working gRPC server
	s.grpcps = grpc_proxy_server.New(cfg.Port)
	if err != nil {
		return nil, fmt.Errorf("couldn't create gRPC proxy server. " + err.Error())
	}

	// Run gRPC server before creating gRPC proxy to allow gRPC proxy dial connection with gRPC
	if err := s.grpcs.Run(s.grpcListener); err != nil {
		return nil, fmt.Errorf("couldn't run gRPC server. " + err.Error())
	}

	if err := s.grpcps.RegisterServices(s.services); err != nil {
		return nil, fmt.Errorf("couldn't register gRPC proxy services. " + err.Error())
	}

	sws, err := swagger.New(cfg.Docs.DocsFS, cfg.Docs.DocsRootFolder)
	if err != nil {
		return nil, fmt.Errorf("couldn't create Swagger docs server. " + err.Error())
	}

	s.httpServer = http.New(s.grpcps, sws)

	return s, nil
}

func (s *Server) Run() error {
	s.httpServer.Run(s.httpListener)

	s.log.Info("Wait exit signal")
	quitSignal := make(chan os.Signal, 1)
	signal.Notify(quitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-quitSignal

	s.log.Info("Closing listeners")
	if err := s.listener.Close(); err != nil {
		return err
	}

	return nil
}
