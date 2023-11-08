package config

import "github.com/gobuffalo/envy"

type Config struct {
	HttpServer HttpServer
	DB         Database
}

type HttpServer struct {
	Host string
	Port string
}
type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
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
		DB: Database{
			Host:     envy.Get("POSTGRES_HOST", "localhost"),
			Port:     envy.Get("POSTGRES_PORT", "5432"),
			User:     envy.Get("POSTGRES_USER", "postgres"),
			Password: envy.Get("POSTGRES_PASSWORD", "postgres"),
			Name:     envy.Get("POSTGRES_NAME", "postgres"),
		},
	}
}
