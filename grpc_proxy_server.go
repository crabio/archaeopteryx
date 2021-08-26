package archaeopteryx

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
	api_data "github.com/iakrevetkho/archaeopteryx/pkg/api/data"
	api_health_v1 "github.com/iakrevetkho/archaeopteryx/pkg/api/health/v1"
	"github.com/iakrevetkho/archaeopteryx/pkg/helpers"
)

type internalGrpcProxyServiceRegistrar func(router *runtime.ServeMux, conn *grpc.ClientConn, controllers *api_data.Controllers) error
type ExternalGrpcProxyServiceRegistrar func(router *runtime.ServeMux, conn *grpc.ClientConn, externalControllers interface{}) error

var (
	internalGrpcProxyServicesRegistrars = []internalGrpcProxyServiceRegistrar{
		api_health_v1.RegisterProxyServiceServer,
	}
)

type grpcProxyServer struct {
	log        *logrus.Entry
	port       int
	grpcServer *grpcServer

	grpcConn   *grpc.ClientConn
	httpServer *http.Server
}

// Function creates gRPC server proxy
// to process REST HTTP requests on the [port]
// and proxy them onto gRPC server on [grpcServer] port.
//
// Requests from the [port] will be redirected to the [grpcServer] port.
func newGrpcProxyServer(port int, grpcServer *grpcServer, controllers *api_data.Controllers, externalServicesRegistrars []ExternalGrpcProxyServiceRegistrar, externalControllers interface{}) (*grpcProxyServer, error) {
	s := new(grpcProxyServer)
	s.log = helpers.CreateComponentLogger("grpc-proxy")
	s.port = port
	s.grpcServer = grpcServer

	// Create a client connection to the gRPC server
	var err error
	s.grpcConn, err = grpc.DialContext(
		context.Background(),
		":"+strconv.Itoa(s.grpcServer.port),
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

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

	// Register internal proxy service routes
	for _, servicesRegistrar := range internalGrpcProxyServicesRegistrars {
		if err := servicesRegistrar(mux, s.grpcConn, controllers); err != nil {
			return nil, err
		}
	}
	s.log.Debug("Internal services are registered")
	// Register internal proxy service routes
	for _, servicesRegistrar := range externalServicesRegistrars {
		if err := servicesRegistrar(mux, s.grpcConn, controllers); err != nil {
			return nil, err
		}
	}
	s.log.Debug("External services are registered")

	s.httpServer = &http.Server{
		Addr:    ":" + strconv.Itoa(s.port),
		Handler: mux,
	}

	return s, nil
}

// Function runs gRPC proxy server on the [port]
func (ps *grpcProxyServer) run() error {
	go func() {
		if err := ps.httpServer.ListenAndServe(); err != nil {
			ps.log.WithError(err).Fatal("Couldn't serve gRPC-Gateway server")
		}
	}()

	ps.log.WithField("url", "http://0.0.0.0:"+strconv.Itoa(ps.port)).Debug("Serving gRPC-Gateway")
	return nil
}
