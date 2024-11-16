package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"tenant-management-service/internal/config"
	"tenant-management-service/internal/model"
)

func Connect(config config.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User, config.Password, config.Host, config.Port, config.DBName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	log.Println("Database connection established")
	return db, nil
}

func Close(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("Error closing database connection:", err)
		return
	}
	err = sqlDB.Close()
	if err != nil {
		return
	}
}

func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.Tenant{},
		&model.Configuration{},
		&model.Quota{},
		&model.Usage{},
	)
}
