package http_test

import (
	// External

	"testing"
	"time"

	"github.com/jinzhu/configor"
	"github.com/stretchr/testify/assert"

	// Internal
	"github.com/iakrevetkho/archaeopteryx/config"
	api_health_v1 "github.com/iakrevetkho/archaeopteryx/pkg/api/health/v1"
	"github.com/iakrevetkho/archaeopteryx/pkg/grpc_proxy_server"
	"github.com/iakrevetkho/archaeopteryx/pkg/grpc_server"
	"github.com/iakrevetkho/archaeopteryx/pkg/healthchecker"
	"github.com/iakrevetkho/archaeopteryx/pkg/http"
	"github.com/iakrevetkho/archaeopteryx/pkg/swagger"
	"github.com/iakrevetkho/archaeopteryx/service"
)

func TestNew(t *testing.T) {
	cfg := new(config.Config)
	assert.NoError(t, configor.Load(cfg))

	hc := healthchecker.New()

	services := []service.IServiceServer{api_health_v1.New(hc, time.Second*10)}

	grpcs, err := grpc_server.New(cfg.GrpcPort, services)
	assert.NoError(t, err)
	assert.NotNil(t, grpcs)

	grpcps := grpc_proxy_server.New(cfg.GrpcPort, cfg.RestApiPort)
	assert.NotNil(t, grpcps)

	sws, err := swagger.New(nil, cfg.Docs.DocsRootFolder)
	assert.NoError(t, err)

	httpServer := http.New(cfg.RestApiPort, grpcps, sws)
	assert.NotNil(t, httpServer)
}
