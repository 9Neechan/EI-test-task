package config

import (
	"github.com/spf13/viper"
)

// Config represents the application configuration.
type Config struct {
	DBDriver          string `mapstructure:"DB_DRIVER"`
	DBSource          string `mapstructure:"DB_SOURCE"`
	GRPCServerAddress string `mapstructure:"GRPC_SERVER_ADDRESS"`
}

// LoadConfig reads configuration from a file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path + "/configs/")
	viper.SetConfigName("cfg")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
