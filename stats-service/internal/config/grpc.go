package config

import (
	"os"

	"github.com/pkg/errors"
)

// grpcAdressEnvName is the environment variable name for the GRPC server address
const grpcAdressEnvName = "GRPC_SERVER_ADDRESS"

// GRPCConfig is an interface for accessing GRPC server configuration
type GRPCConfig interface {
	Address() string
}

// grpcConfig holds the configuration for the GRPC server
type grpcConfig struct {
	adress string
}

// NewGRPCConfig creates a new GRPC configuration from environment variables
func NewGRPCConfig() (GRPCConfig, error) {
	adress := os.Getenv(grpcAdressEnvName)
	if len(adress) == 0 {
		return nil, errors.New("grpc server address not found")
	}

	return &grpcConfig{
		adress: adress,
	}, nil
}

// Address returns the address of the GRPC server
func (cfg *grpcConfig) Address() string {
	return cfg.adress
}
