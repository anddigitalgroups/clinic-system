package main

import (
	"log"

	"clinic-system/internal/api"
	"clinic-system/internal/database"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.NewDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	defer db.Close()

	log.Println("Database connected successfully")

	router := api.SetupRouter()

	router.Run(":8080")
}
