package repository

import (
	"app/internal/domain/model"
	"app/internal/domain/repository"

	"gorm.io/gorm"
)

type GormUserRepository struct {
	repository.GenericRepository[model.User, repository.CreateUserInput]
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{
		GenericRepository: NewGenericRepository(
			db,
			func(cmd repository.CreateUserInput) *model.User {
				return &model.User{
					Firstname:      cmd.Firstname,
					Lastname:       cmd.Lastname,
					Email:          cmd.Email,
					HashedPassword: cmd.HashedPassword,
					IsActive:       true,
				}
			},
		),
	}
}

func (r *GormUserRepository) UpdateUserPassword(id uint, updates repository.UpdateUserPasswordInput) error {
	return r.GenericRepository.UpdateByID(id, updates)
}
