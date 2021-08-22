package api

import (
	// External
	"context"
	"net/http"
	"strconv"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"

	// Internal
	api_hello_world_v1 "github.com/iakrevetkho/archaeopteryx/pkg/api/hello_world/v1"
	api_user_v1 "github.com/iakrevetkho/archaeopteryx/pkg/api/user/v1"
	api_user_v2 "github.com/iakrevetkho/archaeopteryx/pkg/api/user/v2"
	"github.com/iakrevetkho/archaeopteryx/pkg/helpers"
)

var (
	// List with all proxy service registars.
	// If you add new service, you need to add registrar here.
	grpcProxyServicesRegistrars = []func(*runtime.ServeMux, *grpc.ClientConn) error{
		api_hello_world_v1.RegisterProxyServiceServer,
		api_user_v1.RegisterProxyServiceServer,
		api_user_v2.RegisterProxyServiceServer,
	}
)

type grpcProxyServer struct {
	log        *logrus.Entry
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
	server.log = helpers.CreateComponentLogger("grpc")
	server.port = port
	server.grpcServer = grpcServer

	// Create mux router to route HTTP requests in server
	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(
			runtime.MIMEWildcard,
			&runtime.JSONPb{
				MarshalOptions: protojson.MarshalOptions{
					// Not use proto names when decode message in request.
					// We need this to have camel case in request and response
					// in JSON field names.
					UseProtoNames: false,
				},
			},
		),
	)

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
	for _, proxyServiceRegistrar := range grpcProxyServicesRegistrars {
		if err := proxyServiceRegistrar(mux, server.grpcConn); err != nil {
			return nil, err
		}
	}
	server.log.Debug("Services are registered")

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
			ps.log.WithError(err).Fatal("Couldn't serve gRPC-Gateway server")
		}
	}()

	ps.log.WithField("url", "http://0.0.0.0:"+strconv.Itoa(ps.port)).Debug("Serving gRPC-Gateway")
	return nil
}
