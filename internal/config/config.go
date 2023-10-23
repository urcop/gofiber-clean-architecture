package config

import "github.com/gobuffalo/envy"

type Config struct {
	HttpServer HttpServer
}

type HttpServer struct {
	Host string
	Port string
}

var config *Config

func GetConfig() *Config {
	if config != nil {
		return config
	}

	return &Config{
		HttpServer: HttpServer{
			Host: envy.Get("HTTP_HOST", "localhost"),
			Port: envy.Get("HTTP_PORT", "8000"),
		},
	}
}
