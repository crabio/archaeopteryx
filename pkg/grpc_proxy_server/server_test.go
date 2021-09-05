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
	"github.com/iakrevetkho/archaeopteryx/pkg/grpc_server"
	"github.com/iakrevetkho/archaeopteryx/service"
)

func TestNew(t *testing.T) {
	c := new(api_data.Controllers)
	c.Config = new(config.Config)
	assert.NoError(t, configor.Load(c.Config))

	grpcs, err := grpc_server.New(c, []service.IServiceServer{})
	assert.NoError(t, err)
	assert.NotNil(t, grpcs)

	// TODO Run serving gRPC
	// assert.NoError(t, grpcs.Run())

	s, err := grpc_proxy_server.New(c.Config, []service.IServiceServer{})
	assert.NoError(t, err)
	assert.NotNil(t, s)
}
