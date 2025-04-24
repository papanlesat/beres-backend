package main

import (
	"beres/config"
	"beres/infra/database"
	"beres/infra/logger"
	"beres/migrations"
	"beres/routers"
	"time"

	"github.com/spf13/viper"
)

func main() {
	// timezone setup as before
	viper.SetDefault("SERVER_TIMEZONE", "Asia/Dhaka")
	loc, _ := time.LoadLocation(viper.GetString("SERVER_TIMEZONE"))
	time.Local = loc

	if err := config.SetupConfig(); err != nil {
		logger.Fatalf("config SetupConfig() error: %s", err)
	}

	// only one DSN now
	dsn := config.DbConfiguration()
	if err := database.DbConnection(dsn); err != nil {
		logger.Fatalf("database DbConnection error: %s", err)
	}

	migrations.Migrate()
	router := routers.SetupRoute()
	logger.Fatalf("%v", router.Run(config.ServerConfig()))
}
