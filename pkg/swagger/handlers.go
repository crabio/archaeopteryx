package swagger

import (
	// External
	"net/http"
	"text/template"

	// Internal
	docs_swagger "github.com/iakrevetkho/archaeopteryx/docs/swagger"
)

func (s *Server) createFsHandler() http.Handler {
	return http.FileServer(http.Dir(s.config.Docs.SwaggerDir))
}

func (s *Server) createHomePageHandler() (http.Handler, error) {
	swaggerFilePaths := []string{
		"health/v1/health_v1.swagger.json",
	}
	swaggerHomeTmpl, err := template.ParseFS(docs_swagger.SwaggerTmpl, "*")
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
