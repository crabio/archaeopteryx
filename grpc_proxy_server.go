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

	"github.com/iakrevetkho/archaeopteryx/pkg/helpers"
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
func newGrpcProxyServer(port int, grpcServer *grpcServer, services []IServiceServer) (*grpcProxyServer, error) {
	s := new(grpcProxyServer)
	s.log = helpers.CreateComponentLogger("archeaopteryx-grpc-proxy")
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
	for _, service := range services {
		if err := service.RegisterGrpcProxy(context.Background(), mux, s.grpcConn); err != nil {
			return nil, err
		}
	}
	s.log.Debug("Services are registered")

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
