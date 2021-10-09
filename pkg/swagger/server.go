package swagger

import (
	// External
	"embed"
	"mime"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	// Internal

	"github.com/iakrevetkho/archaeopteryx/pkg/helpers"
)

type Server struct {
	log *logrus.Entry

	pkgFsHandler  http.Handler
	userFsHandler http.Handler
	hpHandler     http.Handler
}

// Function creates Open API server
func New(docsFS *embed.FS, docsRootFolder string) (*Server, error) {
	var err error
	s := new(Server)
	s.log = helpers.CreateComponentLogger("swagger")

	if err := mime.AddExtensionType(".svg", "image/svg+xml"); err != nil {
		return nil, err
	}

	if docsFS != nil {
		s.pkgFsHandler, err = createFsHandler(docsFS, "swagger")
		if err != nil {
			return nil, err
		}
		s.userFsHandler, err = createFsHandler(docsFS, docsRootFolder)
		if err != nil {
			return nil, err
		}
		s.hpHandler, err = s.createHomePageHandler(docsFS, docsRootFolder)
		if err != nil {
			return nil, err
		}
	} else {
		s.log.Warn("No swagger doc files found")
	}

	return s, nil
}

func (s *Server) GetMainPageHandler() gin.HandlerFunc {
	return gin.WrapH(s.hpHandler)
}

func (s *Server) GetPkgDocsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Change file URL path for FS route
		c.Request.URL.Path = c.Param("file")
		s.pkgFsHandler.ServeHTTP(c.Writer, c.Request)
	}
}

func (s *Server) GetUserDocsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Change file URL path for FS route
		c.Request.URL.Path = c.Param("file")
		s.userFsHandler.ServeHTTP(c.Writer, c.Request)
	}
}
