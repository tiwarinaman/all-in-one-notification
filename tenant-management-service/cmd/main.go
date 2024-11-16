package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"tenant-management-service/internal/api/v1"
	"tenant-management-service/internal/config"
	"tenant-management-service/pkg/database"
	"tenant-management-service/pkg/logger"
)

func main() {

	// Initialize the logger
	logger.InitLogger()
	defer logger.Sync()

	// Load application configs
	appConfig, err := config.LoadConfig("config.yml")
	if err != nil {
		logger.Error("Failed to load configuration", zap.String("error", err.Error()))
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// db connection
	db, err := database.Connect(appConfig.Database)
	if err != nil {
		logger.Error("Failed to connect to database", zap.String("error", err.Error()))
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// close connection
	defer database.Close(db)

	// Run auto migrations (not recommended in prod)
	if err := database.RunMigrations(db); err != nil {
		logger.Error("Failed to run migrations", zap.String("error", err.Error()))
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Initialize Gin router
	router := gin.Default()

	// Register routes
	v1.RegisterRoutes(router, db)

	// Start server
	logger.Info("Server is starting", zap.String("port", appConfig.Server.Port))
	if err := router.Run(appConfig.Server.Port); err != nil {
		logger.Error("Failed to start server", zap.String("error", err.Error()))
		log.Fatalf("Failed to start server: %v", err)
	}

}
