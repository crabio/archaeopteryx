package http

import (
	// External

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/toorop/gin-logrus"

	// Internal
	"github.com/iakrevetkho/archaeopteryx/config"
	"github.com/iakrevetkho/archaeopteryx/pkg/grpc_proxy_server"
	"github.com/iakrevetkho/archaeopteryx/pkg/helpers"
	"github.com/iakrevetkho/archaeopteryx/pkg/swagger"
)

type Server struct {
	c   *config.Config
	log *logrus.Entry
	r   *gin.Engine
}

func New(c *config.Config, grpcps *grpc_proxy_server.Server, sws *swagger.Server) *Server {
	s := new(Server)
	s.log = helpers.CreateComponentLogger("archeaopteryx-http")
	s.c = c
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
	go func() {
		if err := s.r.Run(":" + strconv.FormatUint(s.c.RestApiPort, 10)); err != nil {
			s.log.WithError(err).Fatal("Couldn't run http server")
		}
	}()
}
