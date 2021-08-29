package api_openapi_v1

import (
	// External

	"context"
	"io/fs"
	"mime"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	// Internal
	api_data "github.com/iakrevetkho/archaeopteryx/pkg/api/data"
	"github.com/iakrevetkho/archaeopteryx/pkg/helpers"
	"github.com/iakrevetkho/archaeopteryx/third_party"
)

// OpenApiServiceServer - service for getting OpenAPI docs.
type OpenApiServiceServer struct {
	log         *logrus.Entry
	controllers *api_data.Controllers
}

// New - creates new instance of OpenApiServiceServer
func New(controllers *api_data.Controllers) *OpenApiServiceServer {
	server := new(OpenApiServiceServer)
	server.log = helpers.CreateComponentLogger("grpc-openapi-v1")
	server.controllers = controllers
	return server
}

// RegisterGrpc - OpenApiServiceServer's method to registrate gRPC service server handlers
func (s *OpenApiServiceServer) RegisterGrpc(sr grpc.ServiceRegistrar) error {
	// Do nothing
	return nil
}

// RegisterGrpcProxy - OpenApiServiceServer's method to registrate gRPC proxy service server handlers
func (s *OpenApiServiceServer) RegisterGrpcProxy(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	mime.AddExtensionType(".svg", "image/svg+xml")
	// Use subdirectory in embedded files
	subFS, err := fs.Sub(third_party.OpenAPI, "OpenAPI")
	if err != nil {
		panic("couldn't create sub filesystem: " + err.Error())
	}

	return mux.HandlePath("GET", "/swagger", func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		http.FileServer(http.FS(subFS)).ServeHTTP(w, req)
	})
}
