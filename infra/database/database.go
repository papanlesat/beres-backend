package database

import (
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// DbConnection opens a single MySQL connection.
func DbConnection(dsn string) error {
	logMode := viper.GetBool("DB_LOG_MODE")
	logLevel := logger.Silent
	if logMode {
		logLevel = logger.Info
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
		return err
	}

	DB = db
	return nil
}

// GetDB returns the singleton *gorm.DB
func GetDB() *gorm.DB {
	return DB
}
