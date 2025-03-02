package config

import (
	"os"

	"github.com/pkg/errors"
)

// grpcAdressEnvName is the environment variable name for the GRPC server address
const grpcAdressEnvName = "GRPC_SERVER_ADDRESS"

// GRPCConfig is an interface for GRPC configuration
type GRPCConfig interface {
	Address() string
}

// grpcConfig is a struct for GRPC configuration
type grpcConfig struct {
	gadress string
}

// NewGRPCConfig creates a new GRPC configuration
func NewGRPCConfig() (GRPCConfig, error) {
	gadress := os.Getenv(grpcAdressEnvName)
	if len(gadress) == 0 {
		return nil, errors.New("grpc server address not found")
	}

	return &grpcConfig{
		gadress: gadress,
	}, nil
}

// Address returns the GRPC server address
func (cfg *grpcConfig) Address() string {
	return cfg.gadress
}
