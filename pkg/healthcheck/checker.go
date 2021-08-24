package healthcheck

import (
	// External

	"time"

	"github.com/alexliesenfeld/health"
)

func New() *health.Checker {
	// Create a new Checker.
	checker := health.NewChecker(

		// Set the time-to-live for our cache to 1 second (default).
		health.WithCacheDuration(1*time.Second),

		// Configure a global timeout that will be applied to all checks.
		health.WithTimeout(10*time.Second),
	)
	return &checker
}
