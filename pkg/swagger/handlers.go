package swagger

import (
	// External
	"net/http"
	"text/template"

	"github.com/iakrevetkho/archaeopteryx/docs"
	// Internal
)

func (s *Server) createFsHandler() http.Handler {
	return http.FileServer(http.Dir(s.config.Docs.SwaggerDir))
}

func (s *Server) createHomePageHandler() (http.Handler, error) {
	swaggerFilePaths := []string{
		"health/v1/health_v1.swagger.json",
	}
	swaggerHomeTmpl, err := template.ParseFS(docs.Swagger, "index.html")
	if err != nil {
		return nil, err
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Open API home page
		if err := swaggerHomeTmpl.Execute(w, map[string]interface{}{"filePaths": swaggerFilePaths}); err != nil {
			s.log.WithError(err).Error("couldn't execute OpenAPI home template")
		}
	}), nil
}
