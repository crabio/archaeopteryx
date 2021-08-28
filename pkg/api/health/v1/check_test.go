package api_health_v1_test

import (
	// External
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	// Internal
	api_data "github.com/iakrevetkho/archaeopteryx/pkg/api/data"
	api_health_v1 "github.com/iakrevetkho/archaeopteryx/pkg/api/health/v1"
	"github.com/iakrevetkho/archaeopteryx/pkg/healthchecker"
	health_v1 "github.com/iakrevetkho/archaeopteryx/proto/gen/health/v1"
)

func TestCheck(t *testing.T) {
	c := new(api_data.Controllers)
	c.HealthChecker = healthchecker.New()
	s := api_health_v1.New(c)

	requests := []health_v1.HealthCheckRequest{
		{},
		{Service: ""},
		{Service: "main"},
	}

	for _, request := range requests {
		response, err := s.Check(context.Background(), &request)

		assert.NoError(t, err)
		assert.Equal(t, response.Status, health_v1.HealthCheckResponse_SERVING)
	}
}
