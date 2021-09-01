package swagger_test

import (
	// External

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jinzhu/configor"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	// Internal
	"github.com/iakrevetkho/archaeopteryx/config"
	"github.com/iakrevetkho/archaeopteryx/docs"
	"github.com/iakrevetkho/archaeopteryx/pkg/helpers"
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

	req, err := http.NewRequest("GET", "/doc/swagger", nil)
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
	conf.Log.Level = logrus.DebugLevel
	helpers.InitLogger(conf)

	s, err := swagger.New(conf)
	assert.NoError(t, err)
	assert.NotNil(t, s)

	req, err := http.NewRequest("GET", "/doc/swagger/swagger-ui-bundle.js", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	s.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestServePkgDocs(t *testing.T) {
	conf := new(config.Config)
	assert.NoError(t, configor.Load(conf))
	conf.Log.Level = logrus.DebugLevel
	helpers.InitLogger(conf)

	s, err := swagger.New(conf)
	assert.NoError(t, err)
	assert.NotNil(t, s)

	req, err := http.NewRequest("GET", "/doc/swagger/health/v1/health_v1.swagger.json", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	s.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
func TestServeUsrDocs(t *testing.T) {
	conf := new(config.Config)
	assert.NoError(t, configor.Load(conf))
	conf.Log.Level = logrus.DebugLevel
	helpers.InitLogger(conf)
	conf.Docs.DocsFS = &docs.Swagger
	conf.Docs.DocsRootFolder = "swagger"

	s, err := swagger.New(conf)
	assert.NoError(t, err)
	assert.NotNil(t, s)

	req, err := http.NewRequest("GET", "/doc/health/v1/health_v1.swagger.json", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	s.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
