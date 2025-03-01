package config

import (
	"os"

	"github.com/pkg/errors"
)

const (
	dbDriverEnvName = "DB_DRIVER"
	dbSourceEnvName = "DB_SOURCE"
)

type DBConfig interface {
	Driver() string
	Source() string
}

type dbConfig struct {
	driver string
	source string
}

func NewDBConfig() (DBConfig, error) {
	driver := os.Getenv(dbDriverEnvName)
	if len(driver) == 0 {
		return nil, errors.New("db driver not found")
	}

	source := os.Getenv(dbSourceEnvName)
	if len(source) == 0 {
		return nil, errors.New("db sourece not found")
	}

	return &dbConfig{
		driver: driver,
		source: source,
	}, nil
}

func (cfg *dbConfig) Driver() string {
	return cfg.driver
}

func (cfg *dbConfig) Source() string {
	return cfg.source
}
