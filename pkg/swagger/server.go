package swagger

import (
	// External
	"mime"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	// Internal
	"github.com/iakrevetkho/archaeopteryx/config"
	"github.com/iakrevetkho/archaeopteryx/docs"
	"github.com/iakrevetkho/archaeopteryx/pkg/helpers"
)

type Server struct {
	log    *logrus.Entry
	config *config.Config

	pkgFsHandler  http.Handler
	userFsHandler http.Handler
	hpHandler     http.Handler
}

// Function creates Open API server
func New(config *config.Config) (*Server, error) {
	var err error
	s := new(Server)
	s.log = helpers.CreateComponentLogger("swagger")
	s.config = config

	if err := mime.AddExtensionType(".svg", "image/svg+xml"); err != nil {
		return nil, err
	}

	s.pkgFsHandler, err = createFsHandler(docs.Swagger, "swagger")
	if err != nil {
		return nil, err
	}
	if s.config.Docs.DocsFS != nil {
		s.userFsHandler, err = createFsHandler(*s.config.Docs.DocsFS, s.config.Docs.DocsRootFolder)
		if err != nil {
			return nil, err
		}
	}
	s.hpHandler, err = s.createHomePageHandler()
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Server) GetMainPageHandler() gin.HandlerFunc {
	return gin.WrapH(s.hpHandler)
}

func (s *Server) GetPkgDocsHandler() gin.HandlerFunc {
	return gin.WrapH(s.pkgFsHandler)
}

func (s *Server) GetUserDocsHandler() gin.HandlerFunc {
	return gin.WrapH(s.userFsHandler)
}
