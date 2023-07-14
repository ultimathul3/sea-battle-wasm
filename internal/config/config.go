package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

const (
	EnvFilePath = ".env"
)

type Config struct {
	HttpServer struct {
		Host string `env:"HTTP_SERVER_HOST"`
		Port uint16 `env:"HTTP_SERVER_PORT"`
	}
}

func ReadEnvFile() (*Config, error) {
	var cfg Config

	err := cleanenv.ReadConfig(EnvFilePath, &cfg)

	return &cfg, err
}
