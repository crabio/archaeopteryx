package config

import (
	// External
	"github.com/jinzhu/configor"
	"github.com/sirupsen/logrus"
)

const CONFIG_FILE_PATH = "config.yml"

type Config struct {
	Version string `default:"" env:"VERSION"`

	Log struct {
		Level logrus.Level `default:"info" env:"LOG_LEVEL"`
	}
}

// Load config from environment variables
// or from [CONFIG_FILE_PATH] file
// or use default values
func LoadConfig() *Config {
	config := new(Config)
	configor.Load(&config, CONFIG_FILE_PATH)
	return config
}
