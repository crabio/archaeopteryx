package api_health_v1_test

import (
	// External
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	// Internal
	api_health_v1 "github.com/iakrevetkho/archaeopteryx/pkg/api/health/v1"
	"github.com/iakrevetkho/archaeopteryx/pkg/healthchecker"
	health_v1 "github.com/iakrevetkho/archaeopteryx/proto/gen/health/v1"
)

func TestCheck(t *testing.T) {
	hc := healthchecker.New()
	s := api_health_v1.New(hc, time.Second*10)

	testCheck(t, s, &health_v1.HealthCheckRequest{})
	testCheck(t, s, &health_v1.HealthCheckRequest{Service: ""})
	testCheck(t, s, &health_v1.HealthCheckRequest{Service: "main"})
}

func testCheck(t *testing.T, s *api_health_v1.HealthServiceServer, r *health_v1.HealthCheckRequest) {
	response, err := s.Check(context.Background(), r)

	assert.NoError(t, err)
	assert.Equal(t, response.Status, health_v1.HealthCheckResponse_SERVING)
}
