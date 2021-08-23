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
		Level    logrus.Level `default:"info" env:"LOG_LEVEL"`
		Filename string       `default:"/var/log/archaeopteryx/archaeopteryx.log" env:"LOG_FILE_NAME"`
		// If log file size is bigger than this threshold, it will be rotated
		MaxSizeInMb int `default:"20" env:"LOG_MAX_SIZE_IN_MB"`
		// If old log file has age more than 30 days, it will be deleted
		MaxAgeInDays int  `default:"30" env:"LOG_MAX_AGE_IN_DAYS"`
		MaxBackups   int  `default:"30" env:"LOG_MAX_BACKUPS"`
		Compress     bool `default:"true" env:"LOG_COMPRESS_OLD_FILES"`
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
