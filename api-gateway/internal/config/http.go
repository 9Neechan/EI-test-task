package config

import (
	"os"

	"github.com/pkg/errors"
)

const (
	httpAdressEnvName = "HTTP_SERVER_ADDRESS"
)

type HTTPConfig interface {
	Address() string
}

type httpConfig struct {
	hadress string
}

func NewHTTPConfig() (HTTPConfig, error) {
	hadress := os.Getenv(httpAdressEnvName)
	if len(hadress) == 0 {
		return nil, errors.New("http server adress not found")
	}

	return &httpConfig{
		hadress: hadress,
	}, nil
}

func (cfg *httpConfig) Address() string {
	return cfg.hadress
}
