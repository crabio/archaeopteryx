package api_health_v1

import (
	// External
	"context"
	// Internal
	health_v1 "github.com/iakrevetkho/archaeopteryx/proto/gen/health/v1"
)

func (s *HealthServiceServer) Check(ctx context.Context, request *health_v1.HealthCheckRequest) (*health_v1.HealthCheckResponse, error) {
	s.log.WithField("request", request.String()).Trace("Check request")

	healthStatus := s.getHealthStatus(ctx)
	response := health_v1.HealthCheckResponse{Status: healthStatus}
	s.log.WithField("response", response.String()).Trace("Check response")

	return &response, nil
}
