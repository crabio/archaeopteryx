package api

import (
	// External
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"

	// Internal
	api_hello_world_v1 "github.com/iakrevetkho/archaeopteryx/pkg/api/hello_world/v1"
	api_user_v1 "github.com/iakrevetkho/archaeopteryx/pkg/api/user/v1"
)

var (
	// List with all service registars.
	// If you add new service, you need to add registrar here.
	grpcServicesRegistrars = []func(grpc.ServiceRegistrar) error{
		api_hello_world_v1.RegisterServiceServer,
		api_user_v1.RegisterServiceServer,
	}
)

type grpcServer struct {
	port       int
	grpcServer *grpc.Server
}

// Function creates gRPC server on the [port]
func newGrpcServer(port int) (*grpcServer, error) {
	server := new(grpcServer)
	server.port = port
	server.grpcServer = grpc.NewServer()

	// Register service routes
	for _, serviceRestrar := range grpcServicesRegistrars {
		if err := serviceRestrar(server.grpcServer); err != nil {
			return nil, err
		}
	}

	return server, nil
}

// Function runs gRPC server on the [port]
func (s *grpcServer) run() error {
	// Create a listener on TCP port
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(s.port))
	if err != nil {
		return err
	}

	go func() {
		if err := s.grpcServer.Serve(listener); err != nil {
			log.Fatalln(err)
		}
	}()
	log.Println("Serving gRPC on 0.0.0.0:" + strconv.Itoa(s.port))

	return nil
}
