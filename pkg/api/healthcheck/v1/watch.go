package api_healthcheck_v1

import (
	// External

	"time"

	// Internal
	healthcheck_v1 "github.com/iakrevetkho/archaeopteryx/proto/gen/healthcheck/v1"
)

func (s *HealthcheckServiceServer) Watch(request *healthcheck_v1.WatchRequest, stream healthcheck_v1.HealthCheckService_WatchServer) error {
	s.log.WithField("request", request.String()).Trace("Watch request")

	// Infinite loop till error on send
	for {
		ctx := stream.Context()
		healthStatus, details := s.getHealthStatus(ctx)
		response := &healthcheck_v1.WatchResponse{
			Status: &healthcheck_v1.HealthCheckResponse{
				Status:  healthStatus,
				Details: details,
			},
		}
		if err := stream.Send(response); err != nil {
			s.log.WithError(err).Trace("Stop watch")
			return err
		}
		s.log.WithField("response", response.String()).Trace("Watch response")

		// Sleep till next check
		time.Sleep(watchUpdatePeriod)
	}
}
