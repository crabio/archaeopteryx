package api_health_v1_test

import (
	// External

	"context"
	"errors"
	"sync"
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

	sync.Mutex
	Close bool

	grpc.ServerStream
}

func (s *MockhWatchServer) Send(m *health_v1.HealthCheckResponse) error {
	s.Lock()
	s.MsgCount += 1
	needClose := s.Close
	s.Unlock()

	if needClose {
		return errors.New("connection is closed")
	}
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

	testWatch(t, s, &health_v1.HealthCheckRequest{})
	testWatch(t, s, &health_v1.HealthCheckRequest{Service: ""})
	testWatch(t, s, &health_v1.HealthCheckRequest{Service: "main"})
}

func testWatch(t *testing.T, s *api_health_v1.HealthServiceServer, r *health_v1.HealthCheckRequest) {
	mockWatchServer := MockhWatchServer{}

	exit := make(chan bool)
	go func() {
		assert.Error(t, s.Watch(r, &mockWatchServer))
		exit <- true
	}()

	time.Sleep(time.Millisecond * 250)

	// Close Mock connection
	mockWatchServer.Lock()
	mockWatchServer.Close = true
	assert.Equal(t, uint32(3), mockWatchServer.MsgCount)
	mockWatchServer.Unlock()

	// Wait exit
	<-exit
}
