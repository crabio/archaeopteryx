package api_health_v1

import (
	// External

	"time"

	// Internal
	health_v1 "github.com/iakrevetkho/archaeopteryx/proto/gen/health/v1"
)

func (s *HealthServiceServer) Watch(request *health_v1.HealthCheckRequest, stream health_v1.Health_WatchServer) error {
	s.log.WithField("request", request.String()).Trace("Watch request")

	// Infinite loop till error on send
	for {
		ctx := stream.Context()
		healthStatus := s.getHealthStatus(ctx)
		response := &health_v1.HealthCheckResponse{Status: healthStatus}
		if err := stream.Send(response); err != nil {
			s.log.WithError(err).Trace("Stop watch")
			return err
		}
		s.log.WithField("response", response.String()).Trace("Watch response")

		// Sleep till next check
		time.Sleep(s.controllers.Config.Health.WatchUpdatePeriod)
	}
}
