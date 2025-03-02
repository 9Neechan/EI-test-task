package config

import (
	"os"

	"github.com/pkg/errors"
)

// httpAdressEnvName is the environment variable name for the HTTP server address
const httpAdressEnvName = "HTTP_SERVER_ADDRESS"

// HTTPConfig is an interface for HTTP server configuration
type HTTPConfig interface {
	Address() string
}

// httpConfig is a struct that implements the HTTPConfig interface
type httpConfig struct {
	hadress string
}

// NewHTTPConfig creates a new HTTPConfig instance
func NewHTTPConfig() (HTTPConfig, error) {
	hadress := os.Getenv(httpAdressEnvName)
	if len(hadress) == 0 {
		return nil, errors.New("http server address not found")
	}

	return &httpConfig{
		hadress: hadress,
	}, nil
}

// Address returns the HTTP server address
func (cfg *httpConfig) Address() string {
	return cfg.hadress
}
