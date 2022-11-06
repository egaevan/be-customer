package config

import (
	"golang-demo/pkg/config"
	"os"
	"time"
)

type ServerConfig struct {
	Port    string
	Debug   bool
	Timeout time.Duration
}

func Server() ServerConfig {
	return ServerConfig{
		Port:    os.Getenv("PORT"),
		Debug:   config.ConvertBool("DEBUG"),
		Timeout: time.Duration(config.ConvertInt("APP_TIMEOUT")) * time.Second,
	}
}
