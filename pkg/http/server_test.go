package http_test

import (
	// External

	"testing"

	"github.com/jinzhu/configor"
	"github.com/stretchr/testify/assert"

	// Internal
	"github.com/iakrevetkho/archaeopteryx/config"
	api_data "github.com/iakrevetkho/archaeopteryx/pkg/api/data"
	api_health_v1 "github.com/iakrevetkho/archaeopteryx/pkg/api/health/v1"
	"github.com/iakrevetkho/archaeopteryx/pkg/grpc_proxy_server"
	"github.com/iakrevetkho/archaeopteryx/pkg/grpc_server"
	"github.com/iakrevetkho/archaeopteryx/pkg/http"
	"github.com/iakrevetkho/archaeopteryx/pkg/swagger"
	"github.com/iakrevetkho/archaeopteryx/service"
)

func TestNew(t *testing.T) {
	c := new(api_data.Controllers)
	c.Config = new(config.Config)
	assert.NoError(t, configor.Load(c.Config))
	services := []service.IServiceServer{api_health_v1.New(c)}

	grpcs, err := grpc_server.New(c.Config, c, services)
	assert.NoError(t, err)
	assert.NotNil(t, grpcs)

	assert.NoError(t, grpcs.Run())

	grpcps := grpc_proxy_server.New(c.Config)
	assert.NotNil(t, grpcps)

	assert.NoError(t, grpcps.RegisterServices(services))

	sws, err := swagger.New(c.Config)
	assert.NoError(t, err)

	httpServer := http.New(c.Config, grpcps, sws)
	httpServer.Run()
}
