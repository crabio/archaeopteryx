package grpc_proxy_server

import (
	// External
	"context"
	"io/fs"
	"mime"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"

	// Internal
	"github.com/iakrevetkho/archaeopteryx/pkg/grpc_server"
	"github.com/iakrevetkho/archaeopteryx/pkg/helpers"
	"github.com/iakrevetkho/archaeopteryx/service"
	"github.com/iakrevetkho/archaeopteryx/third_party"
)

type Server struct {
	log        *logrus.Entry
	Port       int
	grpcServer *grpc_server.Server

	grpcConn   *grpc.ClientConn
	httpServer *http.Server

	openApiFilePaths []string
}

// Function creates gRPC server proxy
// to process REST HTTP requests on the [port]
// and proxy them onto gRPC server on [grpcServer] port.
//
// Requests from the [port] will be redirected to the [grpcServer] port.
func New(port int, grpcServer *grpc_server.Server, services []service.IServiceServer) (*Server, error) {
	var err error
	ps := new(Server)
	ps.log = helpers.CreateComponentLogger("archeaopteryx-grpc-proxy")
	ps.Port = port
	ps.grpcServer = grpcServer
	ps.openApiFilePaths, err = getOpenAPIFilesPaths()
	if err != nil {
		return nil, err
	}

	// Create a client connection to the gRPC server
	ps.grpcConn, err = ps.createGrpcProxyConnection()
	if err != nil {
		return nil, err
	}

	// Create mux router to route HTTP requests in server
	mux := createGrpcProxyMuxServer()

	// Register internal proxy service routes
	for _, service := range services {
		if err := service.RegisterGrpcProxy(context.Background(), mux, ps.grpcConn); err != nil {
			return nil, err
		}
	}
	ps.log.Debug("Services are registered")

	handler, err := ps.getGrpcProxyHandler(mux)
	if err != nil {
		return nil, err
	}

	ps.httpServer = &http.Server{
		Addr:    ":" + strconv.Itoa(ps.Port),
		Handler: handler,
	}

	return ps, nil
}

func (ps *Server) createGrpcProxyConnection() (*grpc.ClientConn, error) {
	return grpc.DialContext(
		context.Background(),
		":"+strconv.Itoa(ps.grpcServer.Port),
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
}

func createGrpcProxyMuxServer() *runtime.ServeMux {
	return runtime.NewServeMux(
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
}

// Function runs gRPC proxy server on the [port]
func (ps *Server) Run() error {
	go func() {
		if err := ps.httpServer.ListenAndServe(); err != nil {
			ps.log.WithError(err).Fatal("Couldn't serve gRPC-Gateway server")
		}
	}()

	ps.log.WithField("url", "http://0.0.0.0:"+strconv.Itoa(ps.Port)).Debug("Serving gRPC-Gateway")
	return nil
}

func (ps *Server) getGrpcProxyHandler(mux *runtime.ServeMux) (http.Handler, error) {
	oa, err := getOpenApiFilesHandler()
	if err != nil {
		return nil, err
	}

	oaHomeTmpl, err := template.ParseFiles("third_party/OpenAPI/index.html")
	if err != nil {
		return nil, err
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api") {
			mux.ServeHTTP(w, r)

		} else if r.URL.Path == "/" {
			// Open API home page
			if err := oaHomeTmpl.Execute(w, map[string]interface{}{"filePaths": ps.openApiFilePaths}); err != nil {
				ps.log.WithError(err).Error("couldn't execute OpenAPI home template")
			}

		} else {
			oa.ServeHTTP(w, r)
		}
	}), nil
}

// getOpenApiFilesHandler serves an OpenAPI UI.
// Adapted from https://github.com/philips/grpc-gateway-example/blob/a269bcb5931ca92be0ceae6130ac27ae89582ecc/cmd/serve.go#L63
func getOpenApiFilesHandler() (http.Handler, error) {
	if err := mime.AddExtensionType(".svg", "image/svg+xml"); err != nil {
		return nil, err
	}

	// Use subdirectory in embedded files
	subFS, err := fs.Sub(third_party.OpenAPI, "OpenAPI")
	if err != nil {
		panic("couldn't create sub filesystem: " + err.Error())
	}

	return http.FileServer(http.FS(subFS)), nil
}
