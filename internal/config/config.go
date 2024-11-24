package config

import (
	"github.com/caarlos0/env/v6"
	"time"
)

type Configuration struct {
	HttpServer
	DataStore

	Env string `env:"ENVIRONMENT" description:"env mode" envDefault:"dev"`
}

type HttpServer struct {
	Address     string        `env:"DEBUG" envDefault:"localhost:8080"`
	Timeout     time.Duration `env:"DEBUG" envDefault:"4s"`
	IdleTimeout time.Duration `env:"DEBUG" envDefault:"60s"`
}

type DataStore struct {
	Path   string `env:"DB_PATH" envDefault:"./url.db"`
	DbName string `env:"DB_NAME" envDefault:"sqlite3"`
}

func FromEnvs() (*Configuration, error) {
	c := Configuration{}

	if err := env.Parse(&c); err != nil {
		return nil, err
	}

	return &c, nil
}
