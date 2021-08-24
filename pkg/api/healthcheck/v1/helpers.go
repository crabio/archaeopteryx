package api_healthcheck_v1

import (
	// External
	"context"

	"github.com/alexliesenfeld/health"
	"github.com/iakrevetkho/archaeopteryx/pkg/helpers"

	// Internal
	healthcheck_v1 "github.com/iakrevetkho/archaeopteryx/proto/healthcheck/v1"
)

func (s *HealthcheckServiceServer) getHealthStatus(ctx context.Context) (healthcheck_v1.HealthCheckResponse_ServingStatus, string) {
	status := s.checker.Check(ctx)

	if status.Status == health.StatusUp {
		return healthcheck_v1.HealthCheckResponse_SERVING_STATUS_SERVING, ""
	} else {
		// Get details
		if status.Details != nil {
			return healthcheck_v1.HealthCheckResponse_SERVING_STATUS_NOT_SERVING, helpers.MustMarshal(status.Details)
		} else {
			return healthcheck_v1.HealthCheckResponse_SERVING_STATUS_NOT_SERVING, ""
		}
	}
}
