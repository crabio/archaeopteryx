package grpc_proxy_server

import (
	// External

	"io/fs"
	"mime"
	"net/http"

	// Internal

	"github.com/iakrevetkho/archaeopteryx/third_party"
)

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
