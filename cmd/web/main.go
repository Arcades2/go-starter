package main

import (
	"fmt"
	"log"

	"app/internal/config"
	"app/internal/infrastructure/http/gin/api"
	infraGorm "app/internal/infrastructure/persistence/gorm"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	config.Load()
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.DBHost,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBPort,
	)
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
		panic(err)
	}

	infraGorm.Migrate(DB)

	router := api.SetupRouter(DB)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
