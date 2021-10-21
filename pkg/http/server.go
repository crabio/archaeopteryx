package http

import (
	// External

	"net"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"

	// Internal

	"github.com/iakrevetkho/archaeopteryx/pkg/grpc_proxy_server"
	"github.com/iakrevetkho/archaeopteryx/pkg/helpers"
	"github.com/iakrevetkho/archaeopteryx/pkg/swagger"
)

type Server struct {
	log *logrus.Entry
	r   *gin.Engine
}

func New(grpcps *grpc_proxy_server.Server, sws *swagger.Server) *Server {
	s := new(Server)
	s.log = helpers.CreateComponentLogger("archeaopteryx-http")
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

func (s *Server) Run(l net.Listener) {
	go func() {

		if err := s.r.RunListener(l); err != nil {
			s.log.WithError(err).Fatal("Couldn't run http server")
		}
	}()
}
