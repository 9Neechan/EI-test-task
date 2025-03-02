package config

import (
	"os"

	"github.com/pkg/errors"
)

const (
	grpcAdressEnvName = "GRPC_SERVER_ADDRESS"
)

type GRPCConfig interface {
	Address() string
}

type grpcConfig struct {
	gadress string
}

func NewGRPCConfig() (GRPCConfig, error) {
	gadress := os.Getenv(grpcAdressEnvName)
	if len(gadress) == 0 {
		return nil, errors.New("grpc server adress not found")
	}

	return &grpcConfig{
		gadress: gadress,
	}, nil
}

func (cfg *grpcConfig) Address() string {
	return cfg.gadress
}
