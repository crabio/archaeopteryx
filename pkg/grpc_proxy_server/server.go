package grpc_proxy_server

import (
	// External
	"context"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/encoding/protojson"

	// Internal
	"github.com/iakrevetkho/archaeopteryx/config"
	"github.com/iakrevetkho/archaeopteryx/pkg/helpers"
	"github.com/iakrevetkho/archaeopteryx/service"
)

type Server struct {
	log      *logrus.Entry
	c        *config.Config
	Port     int
	grpcAddr string
	grpcConn *grpc.ClientConn
	mux      *runtime.ServeMux
}

// Function creates gRPC server proxy
// to process REST HTTP requests on the [port]
// and proxy them onto gRPC server on [grpcServer] port.
//
// Requests from the [port] will be redirected to the [grpcServer] port.
func New(c *config.Config) *Server {
	s := new(Server)
	s.log = helpers.CreateComponentLogger("archeaopteryx-grpc-proxy")
	s.c = c
	s.grpcAddr = ":" + strconv.FormatUint(s.c.GrpcPort, 10)

	// Create mux router to route HTTP requests in server
	s.mux = runtime.NewServeMux(
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

	return s
}

func (s *Server) GetHttpHandler() gin.HandlerFunc {
	return gin.WrapH(s.mux)
}

// Method dials connection to the grpc Server and registers user services.
//
// NOTE. Before this function call, gRPC server should be served.
func (s *Server) RegisterServices(services []service.IServiceServer) error {
	var err error

	context, cancelCtx := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelCtx()

	// Create a client connection to the gRPC server
	if s.c.Secutiry.TlsConfig != nil {
		s.log.WithField("grpcAddr", s.grpcAddr).Info("Create secure connection to gRPC server")
		s.grpcConn, err = grpc.DialContext(
			context,
			s.grpcAddr,
			grpc.WithBlock(),
			grpc.WithTransportCredentials(credentials.NewTLS(s.c.Secutiry.TlsConfig)),
		)
		if err != nil {
			return err
		}
	} else {
		s.log.WithField("grpcAddr", s.grpcAddr).Info("Create insecure connection to gRPC server")
		s.grpcConn, err = grpc.DialContext(
			context,
			s.grpcAddr,
			grpc.WithBlock(),
			grpc.WithInsecure(),
		)
		if err != nil {
			return err
		}
	}

	// Register internal proxy service routes
	for _, service := range services {
		if err := service.RegisterGrpcProxy(context, s.mux, s.grpcConn); err != nil {
			return err
		}
	}
	s.log.Debug("Services are registered")

	return nil
}
