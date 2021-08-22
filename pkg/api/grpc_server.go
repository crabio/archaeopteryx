package api

import (
	// External
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"

	// Internal
	api_hello_world_v1 "github.com/iakrevetkho/archaeopteryx/pkg/api/hello_world/v1"
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
	if err := api_hello_world_v1.RegisterServiceServer(server.grpcServer); err != nil {
		return nil, err
	}

	return server, nil
}

// Function runs gRPC server on the [port]
func (s grpcServer) run() error {
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
