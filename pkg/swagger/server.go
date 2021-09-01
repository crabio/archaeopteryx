package swagger

import (
	// External
	"mime"
	"net/http"
	"strings"

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

const (
	mainPagePath   = "/doc/swagger"
	pkgDocsPrefix  = "/doc/achaeopteryx"
	userDocsPrefix = "/doc"
)

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == mainPagePath {
		s.log.Debug("Serve main page")
		s.hpHandler.ServeHTTP(w, r)

	} else if strings.HasPrefix(r.URL.Path, userDocsPrefix) {
		s.log.WithField("path", r.URL.Path).Debug("Serve user docs")
		if s.userFsHandler != nil {
			s.userFsHandler.ServeHTTP(w, r)
		} else {
			s.log.Error("user fs handler is not inited. Add user swagger docs FS")
		}

	} else if strings.HasPrefix(r.URL.Path, userDocsPrefix) {
		s.log.WithField("path", r.URL.Path).Debug("Serve pkg docs")
		s.pkgFsHandler.ServeHTTP(w, r)
	}
}
