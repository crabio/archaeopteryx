package api_health_v1

import (
	// External
	"context"

	"github.com/alexliesenfeld/health"

	// Internal
	api_data "github.com/iakrevetkho/archaeopteryx/pkg/api/data"
	health_v1 "github.com/iakrevetkho/archaeopteryx/proto/gen/health/v1"
)

func (s *HealthServiceServer) getHealthStatus(ctx context.Context) health_v1.HealthCheckResponse_ServingStatus {
	controllers := ctx.Value("controllers").(api_data.Controllers)
	status := controllers.HealthChecker.Check(ctx)

	if status.Status == health.StatusUp {
		return health_v1.HealthCheckResponse_SERVING
	} else {
		return health_v1.HealthCheckResponse_NOT_SERVING
	}
}
