package http

import (
	// External

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iakrevetkho/archaeopteryx/config"
	"github.com/iakrevetkho/archaeopteryx/pkg/grpc_proxy_server"
	"github.com/iakrevetkho/archaeopteryx/pkg/grpc_server"
	"github.com/iakrevetkho/archaeopteryx/pkg/helpers"
	// Internal
)

type Server struct {
	c *config.Config
	r *gin.Engine
}

func New(c *config.Config, grpcs *grpc_server.Server, grpcps *grpc_proxy_server.Server) *Server {
	s := new(Server)
	s.c = c
	s.r = gin.New()

	s.r.POST(helpers.GRPC_PATH, grpcs.GetHttpHandler())
	s.r.POST(helpers.GRPC_PROXY_PATH, grpcps.GetHttpHandler())

	return s
}

func (s *Server) Run() error {
	return s.r.Run(":" + strconv.FormatUint(s.c.Port, 10))
}
