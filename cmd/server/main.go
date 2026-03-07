package main

import (
	"clinic-system/internal/api"
	"clinic-system/internal/database"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	// Initialize logger
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic("failed to initialize logger")
	}
	defer logger.Sync()

	logger.Info("Starting clinic system")

	// Load environment variables
	err = godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env file")
	}

	// Connect database
	db, err := database.NewDB()
	if err != nil {
		logger.Fatal("Failed to connect to database", zap.Error(err))
	}

	defer db.Close()

	logger.Info("Database connected successfully")

	// Setup router (Dependency Injection)
	router := api.SetupRouter(db)

	logger.Info("Server starting", zap.String("port", "8080"))

	err = router.Run(":8080")
	if err != nil {
		logger.Fatal("Server failed to start", zap.Error(err))
	}
}
