package api_health_v1

import (
	// External
	"context"

	"github.com/alexliesenfeld/health"

	// Internal
	health_v1 "github.com/iakrevetkho/archaeopteryx/proto/gen/health/v1"
)

func (s *HealthServiceServer) getHealthStatus(ctx context.Context) health_v1.HealthCheckResponse_ServingStatus {
	status := s.hc.Check(ctx)

	if status.Status == health.StatusUp {
		return health_v1.HealthCheckResponse_SERVING
	} else {
		return health_v1.HealthCheckResponse_NOT_SERVING
	}
}
