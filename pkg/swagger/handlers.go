package swagger

import (
	// External

	"embed"
	"io/fs"
	"net/http"
	"text/template"

	// Internal
	"github.com/iakrevetkho/archaeopteryx/docs"
	docs_swagger "github.com/iakrevetkho/archaeopteryx/docs/swagger"
)

func createFsHandler(fileSystem embed.FS, folder string) (http.Handler, error) {
	subFS, err := fs.Sub(fileSystem, folder)
	if err != nil {
		return nil, err
	}
	return http.FileServer(http.FS(subFS)), nil
}

func (s *Server) createHomePageHandler() (http.Handler, error) {
	var swaggerFilePaths []string

	// Parse and add pkg's swagger files
	pkgSwaggerFilePaths, err := GetSwaggerFilesPaths(docs.Swagger, "swagger", pkgDocsPrefix)
	if err != nil {
		return nil, err
	}
	swaggerFilePaths = append(swaggerFilePaths, pkgSwaggerFilePaths...)

	// Parse and add user's swagger files
	if s.config.Docs.DocsFS != nil {
		userSwaggerFilePaths, err := GetSwaggerFilesPaths(*s.config.Docs.DocsFS, s.config.Docs.DocsRootFolder, userDocsPrefix)
		if err != nil {
			s.log.WithError(err).Error("No user's swagger files found")
		} else {
			swaggerFilePaths = append(swaggerFilePaths, userSwaggerFilePaths...)
		}
	}

	swaggerHomeTmpl, err := template.ParseFS(docs_swagger.SwaggerTmpl, "*")
	if err != nil {
		return nil, err
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Open API home page
		if err := swaggerHomeTmpl.Execute(w, map[string]interface{}{"filePaths": swaggerFilePaths}); err != nil {
			s.log.WithError(err).Error("couldn't execute Swagger home template")
		}
	}), nil
}
