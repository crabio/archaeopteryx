package grpc_proxy_server

import (
	// External
	"mime"
	"net/http"
	// Internal
)

// getOpenApiFilesHandler serves an OpenAPI UI.
// Adapted from https://github.com/philips/grpc-gateway-example/blob/a269bcb5931ca92be0ceae6130ac27ae89582ecc/cmd/serve.go#L63
func getOpenApiFilesHandler() (http.Handler, error) {
	if err := mime.AddExtensionType(".svg", "image/svg+xml"); err != nil {
		return nil, err
	}

	return http.FileServer(http.Dir("./docs/swagger")), nil
}
