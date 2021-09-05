package grpc_proxy_server

import (
	// External
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"

	// Internal
	"github.com/iakrevetkho/archaeopteryx/config"
	"github.com/iakrevetkho/archaeopteryx/pkg/helpers"
	"github.com/iakrevetkho/archaeopteryx/service"
)

type Server struct {
	log      *logrus.Entry
	Port     int
	grpcConn *grpc.ClientConn
	mux      *runtime.ServeMux
}

// Function creates gRPC server proxy
// to process REST HTTP requests on the [port]
// and proxy them onto gRPC server on [grpcServer] port.
//
// Requests from the [port] will be redirected to the [grpcServer] port.
func New(config *config.Config, services []service.IServiceServer) (*Server, error) {
	var err error
	ps := new(Server)
	ps.log = helpers.CreateComponentLogger("archeaopteryx-grpc-proxy")

	// Create a client connection to the gRPC server
	ps.grpcConn, err = grpc.DialContext(
		context.Background(),
		":"+strconv.Itoa(config.Port)+helpers.GRPC_PATH,
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	// Create mux router to route HTTP requests in server
	ps.mux = runtime.NewServeMux(
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
		if err := service.RegisterGrpcProxy(context.Background(), ps.mux, ps.grpcConn); err != nil {
			return nil, err
		}
	}
	ps.log.Debug("Services are registered")

	return ps, nil
}

func (ps *Server) GetGrpcProxyHandler(mux *runtime.ServeMux) gin.HandlerFunc {
	return gin.WrapH(ps.mux)
}
