package gorm

import (
	"log"

	"app/internal/domain/model"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	log.Println("Running database migrations...")

	err := db.AutoMigrate(
		&model.User{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database migrations completed.")
}
