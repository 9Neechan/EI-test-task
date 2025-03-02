package config

import (
	"os"

	"github.com/pkg/errors"
)

// Constants for environment variable names
const (
	dbDriverEnvName = "DB_DRIVER"
	dbSourceEnvName = "DB_SOURCE"
)

// DBConfig interface defines methods for accessing database configuration
type DBConfig interface {
	Driver() string
	Source() string
}

// dbConfig struct holds database configuration
type dbConfig struct {
	driver string
	source string
}

// NewDBConfig creates a new database configuration from environment variables
func NewDBConfig() (DBConfig, error) {
	driver := os.Getenv(dbDriverEnvName)
	if len(driver) == 0 {
		return nil, errors.New("db driver not found")
	}

	source := os.Getenv(dbSourceEnvName)
	if len(source) == 0 {
		return nil, errors.New("db source not found")
	}

	return &dbConfig{
		driver: driver,
		source: source,
	}, nil
}

// Driver returns the database driver
func (cfg *dbConfig) Driver() string {
	return cfg.driver
}

// Source returns the database source
func (cfg *dbConfig) Source() string {
	return cfg.source
}
