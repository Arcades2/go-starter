// Package gorm provides database migration functionality using GORM. It defines a Migrate function that automatically migrates the database schema based on the defined models. The function logs the migration process and handles any errors that may occur during migration.
package gorm

import (
	"log"

	"app/internal/domain/post"
	"app/internal/domain/user"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	log.Println("Running database migrations...")

	err := db.AutoMigrate(
		&user.User{},
		&post.Post{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database migrations completed.")
}
