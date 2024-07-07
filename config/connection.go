package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenConnection(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.Database.DBHost,
		config.Database.DBUser,
		config.Database.DBPass,
		config.Database.DBName,
		config.Database.DBPort,
		config.Database.SSLModels,
		config.Database.TimeZone,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
