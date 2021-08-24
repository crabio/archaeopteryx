package api_healthcheck_v1

import (
	"context"

	"github.com/alexliesenfeld/health"
	"github.com/iakrevetkho/archaeopteryx/pkg/helpers"
	healthcheck_v1 "github.com/iakrevetkho/archaeopteryx/proto/healthcheck/v1"
)

func (s *HealthcheckServiceServer) SayHello(ctx context.Context, request *healthcheck_v1.HealthCheckRequest) (*healthcheck_v1.HealthCheckResponse, error) {
	s.log.WithField("request", request.String()).Trace("Health check  request")

	healthStatus, details := s.getHealthStatus(ctx)
	response := healthcheck_v1.HealthCheckResponse{
		Status:  healthStatus,
		Details: details,
	}
	s.log.WithField("response", response.String()).Trace("Health check response")

	return &response, nil
}

func (s *HealthcheckServiceServer) getHealthStatus(ctx context.Context) (healthcheck_v1.HealthCheckResponse_ServingStatus, string) {
	status := s.checker.Check(ctx)

	if status.Status == health.StatusUp {
		return healthcheck_v1.HealthCheckResponse_SERVING, ""
	} else {
		// Get details
		if status.Details != nil {
			return healthcheck_v1.HealthCheckResponse_NOT_SERVING, helpers.MustMarshal(status.Details)
		} else {
			return healthcheck_v1.HealthCheckResponse_NOT_SERVING, ""
		}
	}
}
