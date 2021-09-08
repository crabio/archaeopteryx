package http

import (
	// External

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"

	// Internal
	"github.com/iakrevetkho/archaeopteryx/config"
	"github.com/iakrevetkho/archaeopteryx/pkg/grpc_proxy_server"
	"github.com/iakrevetkho/archaeopteryx/pkg/helpers"
	"github.com/iakrevetkho/archaeopteryx/pkg/swagger"
)

type Server struct {
	c    *config.Config
	log  *logrus.Entry
	r    *gin.Engine
	addr string
}

func New(c *config.Config, grpcps *grpc_proxy_server.Server, sws *swagger.Server) *Server {
	s := new(Server)
	s.log = helpers.CreateComponentLogger("archeaopteryx-http")
	s.c = c
	s.addr = ":" + strconv.FormatUint(s.c.RestApiPort, 10)
	s.r = gin.New()
	s.r.Use(ginlogrus.Logger(s.log), gin.Recovery())

	s.r.GET(helpers.API_ROUTE, grpcps.GetHttpHandler())
	s.r.POST(helpers.API_ROUTE, grpcps.GetHttpHandler())
	s.r.GET(helpers.MAIN_SWAGGER_PAGE_ROUTE, sws.GetMainPageHandler())
	s.r.GET(helpers.PKG_STATIC_DOCS_ROUTE, sws.GetPkgDocsHandler())
	s.r.GET(helpers.USER_STATIC_DOCS_ROUTE, sws.GetUserDocsHandler())
	s.log.Debug("Routes are registered")

	return s
}

func (s *Server) Run() {
	httpServer := http.Server{
		Addr:      ":" + strconv.FormatUint(s.c.RestApiPort, 10),
		Handler:   s.r,
		TLSConfig: s.c.Secutiry.TlsConfig}

	go func() {
		if s.c.Secutiry.TlsConfig != nil {
			s.log.WithField("addr", s.addr).Info("Run TLS secure server")
			if err := httpServer.ListenAndServeTLS("", ""); err != nil {
				s.log.WithError(err).Fatal("Couldn't run http server")
			}
		} else {
			s.log.WithField("addr", s.addr).Info("Run insecure server")
			if err := httpServer.ListenAndServe(); err != nil {
				s.log.WithError(err).Fatal("Couldn't run http server")
			}
		}
	}()
}
