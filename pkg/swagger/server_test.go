package swagger_test

import (
	// External

	"testing"

	"github.com/jinzhu/configor"
	"github.com/stretchr/testify/assert"

	// Internal
	"github.com/iakrevetkho/archaeopteryx/config"
	"github.com/iakrevetkho/archaeopteryx/pkg/swagger"
)

func TestNew(t *testing.T) {
	cfg := new(config.Config)
	assert.NoError(t, configor.Load(cfg))

	s, err := swagger.New(nil, cfg.Docs.DocsRootFolder)
	assert.NoError(t, err)
	assert.NotNil(t, s)
}
