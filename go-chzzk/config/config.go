package config

import (
	"github.com/pelletier/go-toml/v2"
	"os"
)

type Config struct {
	DB struct {
		Database string
		URL      string
	}
	Kafka struct {
		URL      string
		ClientID string
	}
}

func NewConfig(path string) *Config {
	c := &Config{}

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	if err := toml.NewDecoder(f).Decode(c); err != nil {
		panic(err)
	}

	return c
}
