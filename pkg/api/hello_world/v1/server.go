package api_hello_world_v1

import (
	// External
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	// Internal
	hello_world_v1 "github.com/iakrevetkho/archaeopteryx/proto/hello_world/v1"
)

type HelloServiceServer struct {
	// Required to have revese compatability
	hello_world_v1.UnimplementedHelloServiceServer
}

func RegisterServiceServer(s grpc.ServiceRegistrar, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	// Attach the Hello service to the server
	hello_world_v1.RegisterHelloServiceServer(s, &HelloServiceServer{})

	// Attach handler to global handler
	if err := hello_world_v1.RegisterHelloServiceHandler(context.Background(), mux, conn); err != nil {
		return err
	}

	return nil
}
