package api

import (
	// External
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	// Internal
	api_hello_world_v1 "github.com/iakrevetkho/archaeopteryx/pkg/api/hello_world/v1"
)

type grpcProxyServer struct {
	port       int
	grpcServer *grpcServer

	grpcConn *grpc.ClientConn
	server   *http.Server
}

// Function creates gRPC server proxy
// to process REST HTTP requests on the [port]
// and proxy them onto gRPC server on [grpcServer] port.
//
// Requests from the [port] will be redirected to the [grpcServer] port.
func newGrpcProxyServer(port int, grpcServer *grpcServer) (*grpcProxyServer, error) {
	server := new(grpcProxyServer)
	server.port = port
	server.grpcServer = grpcServer

	// Create mux router to route HTTP requests in server
	mux := runtime.NewServeMux()

	// Create a client connection to the gRPC server
	var err error
	server.grpcConn, err = grpc.DialContext(
		context.Background(),
		":"+strconv.Itoa(server.grpcServer.port),
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	// Register proxy service routes
	if err := api_hello_world_v1.RegisterProxyServiceServer(mux, server.grpcConn); err != nil {
		return nil, err
	}

	server.server = &http.Server{
		Addr:    ":" + strconv.Itoa(server.port),
		Handler: mux,
	}

	return server, nil
}

// Function runs gRPC proxy server on the [port]
func (ps *grpcProxyServer) run() error {
	go func() {
		if err := ps.server.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:" + strconv.Itoa(ps.port))
	return nil
}
