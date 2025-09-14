package repository

import (
	"app/internal/core/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	GenericRepositoryInterface[model.User]
}

type GormUserRepository struct {
	*GenericRepository[model.User]
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{
		GenericRepository: NewGenericRepository[model.User](db),
	}
}
