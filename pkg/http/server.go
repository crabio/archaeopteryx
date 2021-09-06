package http

import (
	// External

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iakrevetkho/archaeopteryx/config"
	"github.com/iakrevetkho/archaeopteryx/pkg/grpc_proxy_server"
	"github.com/iakrevetkho/archaeopteryx/pkg/grpc_server"
	"github.com/iakrevetkho/archaeopteryx/pkg/helpers"
	"github.com/sirupsen/logrus"
	// Internal
)

type Server struct {
	c   *config.Config
	log *logrus.Entry
	r   *gin.Engine
}

func New(c *config.Config, grpcs *grpc_server.Server, grpcps *grpc_proxy_server.Server) *Server {
	s := new(Server)
	s.log = helpers.CreateComponentLogger("archeaopteryx-http")
	s.c = c
	s.r = gin.New()

	s.r.POST(helpers.GRPC_PATH, grpcs.GetHttpHandler())
	s.r.GET(helpers.GRPC_PATH, grpcs.GetHttpHandler())
	s.r.OPTIONS(helpers.GRPC_PATH, grpcs.GetHttpHandler())
	s.r.POST(helpers.GRPC_PROXY_PATH, grpcps.GetHttpHandler())
	s.log.Debug("Routes are registered")

	return s
}

func (s *Server) Run() {
	go func() {
		if err := s.r.Run(":" + strconv.FormatUint(s.c.Port, 10)); err != nil {
			s.log.WithError(err).Fatal("Couldn't run http server")
		}
	}()
}
