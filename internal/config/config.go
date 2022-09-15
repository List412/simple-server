package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

func NewConfig() (*Config, error) {
	_ = godotenv.Load()

	config := Config{}
	err := env.Parse(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

type Config struct {
	HttpServer HttpServer
}

type HttpServer struct {
	Port string `env:"HTTP_SERVER_PORT,required"`
}
