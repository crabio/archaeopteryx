package api_healthcheck_v1

import (
	// External
	"context"
	// Internal
	healthcheck_v1 "github.com/iakrevetkho/archaeopteryx/proto/gen/healthcheck/v1"
)

func (s *HealthcheckServiceServer) Check(ctx context.Context, request *healthcheck_v1.CheckRequest) (*healthcheck_v1.CheckResponse, error) {
	s.log.WithField("request", request.String()).Trace("Check request")

	healthStatus, details := s.getHealthStatus(ctx)
	response := healthcheck_v1.CheckResponse{
		Status: &healthcheck_v1.HealthCheckResponse{
			Status:  healthStatus,
			Details: details,
		},
	}
	s.log.WithField("response", response.String()).Trace("Check response")

	return &response, nil
}
