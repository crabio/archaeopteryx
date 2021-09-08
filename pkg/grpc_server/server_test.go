package grpc_server_test

import (
	// External

	"crypto/tls"
	"testing"

	"github.com/jinzhu/configor"
	"github.com/stretchr/testify/assert"

	// Internal
	"github.com/iakrevetkho/archaeopteryx/config"
	api_data "github.com/iakrevetkho/archaeopteryx/pkg/api/data"
	api_health_v1 "github.com/iakrevetkho/archaeopteryx/pkg/api/health/v1"
	"github.com/iakrevetkho/archaeopteryx/pkg/grpc_server"
	"github.com/iakrevetkho/archaeopteryx/pkg/helpers"
	"github.com/iakrevetkho/archaeopteryx/service"
)

func TestNewInsecure(t *testing.T) {
	c := new(api_data.Controllers)
	c.Config = new(config.Config)
	assert.NoError(t, configor.Load(c.Config))

	s, err := grpc_server.New(c.Config, c, []service.IServiceServer{api_health_v1.New(c)})
	assert.NoError(t, err)
	assert.NotNil(t, s)
}

func TestNewSecure(t *testing.T) {
	c := new(api_data.Controllers)
	c.Config = new(config.Config)
	assert.NoError(t, configor.Load(c.Config))
	// Create TLS config
	tlsCert, err := tls.X509KeyPair(helpers.MockCertBytes, helpers.MockKeyBytes)
	assert.NoError(t, err)
	c.Config.Secutiry.TlsConfig = &tls.Config{
		Certificates:       []tls.Certificate{tlsCert},
		ClientAuth:         tls.NoClientCert,
		InsecureSkipVerify: true,
	}

	s, err := grpc_server.New(c.Config, c, []service.IServiceServer{api_health_v1.New(c)})
	assert.NoError(t, err)
	assert.NotNil(t, s)
}
