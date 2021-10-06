package grpc_proxy_server_test

import (
	// External

	"testing"

	"github.com/jinzhu/configor"
	"github.com/stretchr/testify/assert"

	// Internal
	"github.com/iakrevetkho/archaeopteryx/config"
	api_data "github.com/iakrevetkho/archaeopteryx/pkg/api/data"
	"github.com/iakrevetkho/archaeopteryx/pkg/grpc_proxy_server"
)

func TestNew(t *testing.T) {
	c := new(api_data.Controllers)
	c.Config = new(config.Config)
	assert.NoError(t, configor.Load(c.Config))

	s := grpc_proxy_server.New(c.Config)
	assert.NotNil(t, s)
}
