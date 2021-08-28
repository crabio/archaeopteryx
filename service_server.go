package archaeopteryx

import (
	// External
	"context"

	// Internal
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// IServiceServer - interface for services servers for archaeopteryx.
//
// All services should implements this interface
// to make possible archaeopteryx registrate services handlers in server
type IServiceServer interface {
	// RegisterGrpc - HealthServiceServer's method to registrate gRPC service server handlers
	RegisterGrpc(sr grpc.ServiceRegistrar) error
	// RegisterGrpcProxy - HealthServiceServer's method to registrate gRPC proxy service server handlers
	RegisterGrpcProxy(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error
}
