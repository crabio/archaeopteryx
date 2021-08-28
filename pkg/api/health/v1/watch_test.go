package api_health_v1_test

import (
	// External

	"context"
	"testing"
	"time"

	"github.com/jinzhu/configor"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"

	// Internal

	"github.com/iakrevetkho/archaeopteryx/config"
	api_data "github.com/iakrevetkho/archaeopteryx/pkg/api/data"
	api_health_v1 "github.com/iakrevetkho/archaeopteryx/pkg/api/health/v1"
	"github.com/iakrevetkho/archaeopteryx/pkg/healthchecker"
	health_v1 "github.com/iakrevetkho/archaeopteryx/proto/gen/health/v1"
)

type MockhWatchServer struct {
	MsgCount uint32
	grpc.ServerStream
}

func (s *MockhWatchServer) Send(m *health_v1.HealthCheckResponse) error {
	s.MsgCount += 1
	return nil
}

func (s *MockhWatchServer) Context() context.Context {
	return context.Background()
}

func TestWatch(t *testing.T) {
	c := new(api_data.Controllers)
	c.Config = new(config.Config)
	assert.NoError(t, configor.Load(c.Config))
	c.HealthChecker = healthchecker.New()
	c.Config.Health.WatchUpdatePeriod = time.Millisecond * 100

	s := api_health_v1.New(c)

	requests := []health_v1.HealthCheckRequest{
		{},
		{Service: ""},
		{Service: "main"},
	}

	for _, request := range requests {
		mockWatchServer := MockhWatchServer{}

		go func() {
			assert.NoError(t, s.Watch(&request, &mockWatchServer))
		}()

		time.Sleep(time.Millisecond * 250)

		assert.Equal(t, uint32(3), mockWatchServer.MsgCount)
	}
}
