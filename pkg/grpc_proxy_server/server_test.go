package grpc_proxy_server_test

import (
	// External

	"testing"

	"github.com/jinzhu/configor"
	"github.com/stretchr/testify/assert"

	// Internal
	"github.com/iakrevetkho/archaeopteryx/config"
	"github.com/iakrevetkho/archaeopteryx/pkg/grpc_proxy_server"
)

func TestNew(t *testing.T) {
	cfg := new(config.Config)
	assert.NoError(t, configor.Load(cfg))

	s := grpc_proxy_server.New(cfg.Port)
	assert.NotNil(t, s)
}
