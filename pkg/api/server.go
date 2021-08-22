package api

import (
	// Internal
	"github.com/iakrevetkho/archaeopteryx/pkg/helpers"
)

const (
	GRPC_SERVER_PORT         = 8080
	GRPC_GATEWAY_SERVER_PORT = 8090
)

func RunServer() {
	log := helpers.CreateComponentLogger("api-server")

	grpcServer, err := newGrpcServer(GRPC_SERVER_PORT)
	if err != nil {
		log.WithError(err).Fatal("Couldn't create gRPC server.")
	}
	if err := grpcServer.run(); err != nil {
		log.WithError(err).Fatal("Couldn't run gRPC server")
	}

	grpcProxyServer, err := newGrpcProxyServer(GRPC_GATEWAY_SERVER_PORT, grpcServer)
	if err != nil {
		log.WithError(err).Fatal("Couldn't create gRPC proxy server")
	}
	if err := grpcProxyServer.run(); err != nil {
		log.WithError(err).Fatal("Couldn't run gRPC proxy server")
	}
}
