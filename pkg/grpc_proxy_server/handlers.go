package grpc_proxy_server

import (
	// External
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	// Internal
)

func (ps *Server) getGrpcProxyHandler(mux *runtime.ServeMux) (http.Handler, error) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api") {
			mux.ServeHTTP(w, r)

		} else {
			ps.openApiServer.ServeHTTP(w, r)
		}
	}), nil
}
