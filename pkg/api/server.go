package api

import (
	// External
	"log"
)

const (
	GRPC_SERVER_PORT         = 8080
	GRPC_GATEWAY_SERVER_PORT = 8090
)

func NewServer() {
	grpcServer, err := newGrpcServer(GRPC_SERVER_PORT)
	if err != nil {
		log.Fatal("Couldn't create gRPC server. " + err.Error())
	}
	if err := grpcServer.run(); err != nil {
		log.Fatal("Couldn't run gRPC server. " + err.Error())
	}

	grpcProxyServer, err := newGrpcProxyServer(GRPC_GATEWAY_SERVER_PORT, grpcServer)
	if err != nil {
		log.Fatal("Couldn't create gRPC proxy server. " + err.Error())
	}
	if err := grpcProxyServer.run(); err != nil {
		log.Fatal("Couldn't run gRPC proxy server. " + err.Error())
	}
}
