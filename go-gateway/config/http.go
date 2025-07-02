package config

import "github.com/kopher1601/go-studies/go-gateway/types/http"

type HttpCfg struct {
	Router  []Router `yaml:"router"`
	BaseURL string   `yaml:"base_url"`
}

type Router struct {
	Method   http.HttpMethod `yaml:"method"`
	GetType  http.GetType    `yaml:"get_type"`
	Variable []string        `yaml:"variable"`
	Path     string          `yaml:"path"`

	Auth   *Auth             `yaml:"auth"`
	Header map[string]string `yaml:"header"`
}

type Auth struct {
	Key   string `yaml:"key"`
	Token string `yaml:"token"`
}
