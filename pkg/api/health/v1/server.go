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
	api_data "github.com/iakrevetkho/archaeopteryx/pkg/api/data"
	"github.com/iakrevetkho/archaeopteryx/pkg/helpers"
	health_v1 "github.com/iakrevetkho/archaeopteryx/proto/gen/health/v1"
)

const (
	watchUpdatePeriod = time.Second * 15
)

// HealthServiceServer - service for checking service health.
// Created bases on https://github.com/grpc/grpc/blob/master/doc/health-checking.md
type HealthServiceServer struct {
	log *logrus.Entry
	// Required to have revese compatability
	health_v1.UnimplementedHealthServer
	checker health.Checker
}

// New - creates new instance of HealthServiceServer
func New(controllers *api_data.Controllers) *HealthServiceServer {
	server := new(HealthServiceServer)
	server.log = helpers.CreateComponentLogger("grpc-healthcheck-v1")
	server.checker = controllers.HealthChecker
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
