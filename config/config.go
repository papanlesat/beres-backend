// config/config.go
package config

import (
	"beres/infra/logger"

	"github.com/spf13/viper"
)

// SetupConfig loads environment variables from .env
func SetupConfig() error {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		logger.Errorf("Error reading config file: %v", err)
		return err
	}
	return nil
}
