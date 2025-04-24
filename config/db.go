package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// DbConfiguration builds a MySQL DSN from env vars.
func DbConfiguration() string {
	user := viper.GetString("DB_USER")
	password := viper.GetString("DB_PASSWORD")
	host := viper.GetString("DB_HOST")
	port := viper.GetString("DB_PORT")
	name := viper.GetString("DB_NAME")

	// add any extra params you need here
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, name,
	)
}
