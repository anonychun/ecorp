package config

import (
	"github.com/anonychun/ecorp/internal/bootstrap"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/samber/do"
)

func init() {
	do.Provide(bootstrap.Injector, NewConfig)
}

type Config struct {
	Server struct {
		Port int `envconfig:"port"`
	} `envconfig:"server"`

	Database struct {
		Sql struct {
			Host     string `envconfig:"host"`
			Port     int    `envconfig:"port"`
			User     string `envconfig:"user"`
			Password string `envconfig:"password"`
			Name     string `envconfig:"name"`
		} `envconfig:"sql"`
	} `envconfig:"database"`
}

func NewConfig(i *do.Injector) (*Config, error) {
	godotenv.Load(".env")

	config := &Config{}
	err := envconfig.Process("", config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
