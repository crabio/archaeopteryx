package swagger_test

import (
	// External

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jinzhu/configor"
	"github.com/stretchr/testify/assert"

	// Internal
	"github.com/iakrevetkho/archaeopteryx/config"
	"github.com/iakrevetkho/archaeopteryx/pkg/swagger"
)

func TestNew(t *testing.T) {
	conf := new(config.Config)
	assert.NoError(t, configor.Load(conf))

	s, err := swagger.New(conf)
	assert.NoError(t, err)
	assert.NotNil(t, s)
}

func TestServeMainPage(t *testing.T) {
	conf := new(config.Config)
	assert.NoError(t, configor.Load(conf))

	s, err := swagger.New(conf)
	assert.NoError(t, err)
	assert.NotNil(t, s)

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	s.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestServeJS(t *testing.T) {
	conf := new(config.Config)
	assert.NoError(t, configor.Load(conf))

	s, err := swagger.New(conf)
	assert.NoError(t, err)
	assert.NotNil(t, s)

	req, err := http.NewRequest("GET", "/swagger-ui-bundle.js", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	s.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
