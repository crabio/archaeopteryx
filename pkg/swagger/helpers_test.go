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
	filePaths, err := swagger.GetOpenAPIFilesPaths(docs.Swagger, "swagger", "/abc/")
	assert.NoError(t, err)
	assert.Equal(t, 3, len(filePaths))
	assert.Equal(t, "/abc/swagger/google/api/annotations.swagger.json", filePaths[0])
	assert.Equal(t, "/abc/swagger/google/api/http.swagger.json", filePaths[1])
	assert.Equal(t, "/abc/swagger/health/v1/health_v1.swagger.json", filePaths[2])
}
