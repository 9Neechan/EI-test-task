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
	adress string
}

func NewGRPCConfig() (GRPCConfig, error) {
	adress := os.Getenv(grpcAdressEnvName)
	if len(adress) == 0 {
		return nil, errors.New("grpc server adress not found")
	}

	return &grpcConfig{
		adress: adress,
	}, nil
}

func (cfg *grpcConfig) Address() string {
	return cfg.adress
}
