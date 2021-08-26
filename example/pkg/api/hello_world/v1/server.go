package api_hello_world_v1

import (
	// External
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	// Internal
	api_data "github.com/iakrevetkho/archaeopteryx/pkg/api/data"
	"github.com/iakrevetkho/archaeopteryx/pkg/helpers"
	hello_world_v1 "github.com/iakrevetkho/archaeopteryx/proto/gen/hello_world/v1"
)

type HelloServiceServer struct {
	log *logrus.Entry
	// Required to have revese compatability
	hello_world_v1.UnimplementedHelloServiceServer
}

func RegisterServiceServer(s grpc.ServiceRegistrar, controllers *api_data.Controllers) error {
	server := new(HelloServiceServer)
	server.log = helpers.CreateComponentLogger("grpc-hello-world-v1")

	// Attach the Hello service to the server
	hello_world_v1.RegisterHelloServiceServer(s, server)
	server.log.Debug("Service registered")

	return nil
}

func RegisterProxyServiceServer(mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	// Attach handler to global handler
	if err := hello_world_v1.RegisterHelloServiceHandler(context.Background(), mux, conn); err != nil {
		return err
	}

	return nil
}
