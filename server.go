package main

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
	"github.com/iakrevetkho/archaeopteryx/service"
	"github.com/sirupsen/logrus"
)

type Server struct {
	Config *config.Config

	log         *logrus.Entry
	controllers *api_data.Controllers

	services []service.IServiceServer
}

func New(config *config.Config, externalServices []service.IServiceServer) *Server {
	s := new(Server)

	helpers.InitLogger(config)
	s.log = helpers.CreateComponentLogger("archeaopteryx-server")
	s.log.WithField("config", helpers.MustMarshal(config)).Info("Config is inited")

	s.Config = config
	s.controllers = new(api_data.Controllers)
	s.controllers.Config = config
	s.controllers.HealthChecker = healthchecker.New()

	// Add internal services
	s.services = append(s.services, api_health_v1.New(s.controllers))

	// Add external services
	s.services = append(s.services, externalServices...)

	return s
}

func (s *Server) Run() error {
	grpcServer, err := grpc_server.New(s.Config.GrpcPort, s.controllers, s.services)
	if err != nil {
		return fmt.Errorf("couldn't create gRPC server. " + err.Error())
	}
	// Run gRPC server before creating gRPC proxy to allow gRPC proxy dial connection with gRPC
	if err := grpcServer.Run(); err != nil {
		return fmt.Errorf("couldn't run gRPC server. " + err.Error())
	}

	grpcProxyServer, err := grpc_proxy_server.New(s.Config, grpcServer, s.services)
	if err != nil {
		return fmt.Errorf("couldn't create gRPC proxy server. " + err.Error())
	}

	if err := grpcProxyServer.Run(); err != nil {
		return fmt.Errorf("couldn't run gRPC proxy server. " + err.Error())
	}

	s.log.Info("Wait exit signal")
	quitSignal := make(chan os.Signal, 1)
	signal.Notify(quitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-quitSignal

	return nil
}
