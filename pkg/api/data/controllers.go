package api_data

import (
	// External
	"github.com/alexliesenfeld/health"
	"github.com/iakrevetkho/archaeopteryx/config"
)

type Controllers struct {
	Config        *config.Config
	HealthChecker health.Checker
}
