package open_api

import (
	// External
	"net/http"
	"text/template"
	// Internal
)

func (s *Server) createFsHandler() http.Handler {
	return http.FileServer(http.Dir(s.config.Docs.SwaggerDir))
}

func (s *Server) createHomePageHandler() (http.Handler, error) {
	filePaths, err := getOpenAPIFilesPaths(s.config.Docs.SwaggerDir)
	if err != nil {
		return nil, err
	}
	oaHomeTmpl, err := template.ParseFiles(s.config.Docs.SwaggerDir + "/index.html")
	if err != nil {
		return nil, err
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Open API home page
		if err := oaHomeTmpl.Execute(w, map[string]interface{}{"filePaths": filePaths}); err != nil {
			s.log.WithError(err).Error("couldn't execute OpenAPI home template")
		}
	}), nil
}
