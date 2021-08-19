package api

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	api_hello_world_v1 "github.com/iakrevetkho/archaeopteryx/pkg/api/hello_world/v1"
)

func NewServer() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()

	// Serve gRPC server
	log.Println("Serving gRPC on 0.0.0.0:8080")
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalln(err)
		}
	}()

	mux := runtime.NewServeMux()

	// Register services
	if err := api_hello_world_v1.RegisterServiceServer(s, mux, conn); err != nil {
		log.Fatalln("Failed to register hello service:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: mux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")

	if err := gwServer.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
