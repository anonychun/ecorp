package config

import (
	"time"

	"github.com/anonychun/benih/internal/bootstrap"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
	"github.com/samber/do/v2"
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

	Storage struct {
		S3 struct {
			Endpoint        string        `envconfig:"endpoint"`
			Bucket          string        `envconfig:"bucket"`
			AccessKeyId     string        `envconfig:"access_key_id"`
			SecretAccessKey string        `envconfig:"secret_access_key"`
			UrlExpiration   time.Duration `envconfig:"url_expiration"`
		} `envconfig:"s3"`
	} `envconfig:"storage"`
}

func NewConfig(i do.Injector) (*Config, error) {
	config := &Config{}
	err := envconfig.Process("", config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
