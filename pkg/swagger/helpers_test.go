package swagger_test

import (
	// External

	"testing"

	"github.com/stretchr/testify/assert"

	// Internal
	"github.com/iakrevetkho/archaeopteryx/docs"
	"github.com/iakrevetkho/archaeopteryx/pkg/swagger"
)

func TestGetOpenAPIFilesPaths(t *testing.T) {
	filePaths, err := swagger.GetOpenAPIFilesPaths(docs.Swagger, "swagger")
	assert.NoError(t, err)
	assert.Equal(t, 3, len(filePaths))
}
