package grpc_server_test

import (
	// External

	"testing"

	"github.com/jinzhu/configor"
	"github.com/stretchr/testify/assert"

	// Internal
	"github.com/iakrevetkho/archaeopteryx/config"
	api_health_v1 "github.com/iakrevetkho/archaeopteryx/pkg/api/health/v1"
	"github.com/iakrevetkho/archaeopteryx/pkg/grpc_server"
	"github.com/iakrevetkho/archaeopteryx/pkg/healthchecker"
	"github.com/iakrevetkho/archaeopteryx/service"
)

func TestNew(t *testing.T) {
	cfg := new(config.Config)
	assert.NoError(t, configor.Load(cfg))

	hc := healthchecker.New()

	s, err := grpc_server.New([]service.IServiceServer{api_health_v1.New(hc, cfg.Health.WatchUpdatePeriod)})
	assert.NoError(t, err)
	assert.NotNil(t, s)
}
