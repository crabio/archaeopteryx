package api_health_v1

import (
	// External

	"context"
	"time"

	"github.com/alexliesenfeld/health"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	// Internal
	"github.com/iakrevetkho/archaeopteryx/pkg/helpers"
	health_v1 "github.com/iakrevetkho/archaeopteryx/proto/gen/health/v1"
)

// HealthServiceServer - service for checking service health.
// Created bases on https://github.com/grpc/grpc/blob/master/doc/health-checking.md
type HealthServiceServer struct {
	// Required to have revese compatability
	health_v1.UnimplementedHealthServer

	log *logrus.Entry
	hc  health.Checker
	// Watch update period
	wup time.Duration
}

// New - creates new instance of HealthServiceServer
func New(hc health.Checker, wup time.Duration) *HealthServiceServer {
	server := new(HealthServiceServer)
	server.log = helpers.CreateComponentLogger("grpc-healthcheck-v1")
	server.hc = hc
	server.wup = wup
	return server
}

// RegisterGrpc - HealthServiceServer's method to registrate gRPC service server handlers
func (s *HealthServiceServer) RegisterGrpc(sr grpc.ServiceRegistrar) error {
	health_v1.RegisterHealthServer(sr, s)
	s.log.Debug("Service registered")
	return nil
}

// RegisterGrpcProxy - HealthServiceServer's method to registrate gRPC proxy service server handlers
func (s *HealthServiceServer) RegisterGrpcProxy(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return health_v1.RegisterHealthHandler(ctx, mux, conn)
}
