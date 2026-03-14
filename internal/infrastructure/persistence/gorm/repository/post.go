package repository

import (
	"app/internal/domain/post"

	"gorm.io/gorm"
)

type GormPostRepository struct {
	*Repository[*post.Post]
}

func NewGormPostRepository(db *gorm.DB) *GormPostRepository {
	return &GormPostRepository{
		NewRepository[*post.Post](db),
	}
}
