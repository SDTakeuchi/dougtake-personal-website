package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type (
	DB struct {
		Database    string `required:"true" envconfig:"DB_NAME"`
		UserName    string `required:"true" envconfig:"DB_USER_NAME"`
		Password    string `required:"true" envconfig:"DB_PASSWORD"`
		Port        string `required:"true" envconfig:"DB_PORT" default:"5432"`
		HostName    string `required:"true" envconfig:"DB_HOST_NAME"`
		SSLMode     string `required:"true" envconfig:"DB_SSL_MODE"`
		Debug       bool   `required:"true" envconfig:"DB_DEBUG" default:"true"`
		ConnTimeout int    `required:"true" envconfig:"DB_CONN_TIMEOUT" default:"9"`
		// 0 means no termination
		MaxOpen int `required:"true" envconfig:"DB_MAX_OPEN" default:"0"`
		MaxIdle int `required:"true" envconfig:"DB_MAX_IDLE" default:"2"`
		// 0 means no termination as long as the process is running
		// set values that can be read by time.ParseDuration()
		MaxLifeTime time.Duration `required:"true" envconfig:"DB_MAX_LIFE_TIME" default:"0"`
	}
	Frontend struct {
		HostName string `required:"true" envconfig:"FRONTEND_HOST_NAME"`
		Port     string `required:"true" envconfig:"FRONTEND_PORT" default:"8000"`
	}
	Config struct {
		DB       DB
		Frontend Frontend
	}
)

var globalConfig Config

func Load() {
	envconfig.MustProcess("", globalConfig)
}

func Get() Config {
	return globalConfig
}
