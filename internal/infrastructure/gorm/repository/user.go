package repository

import (
	"app/internal/domain/model"
	"app/internal/domain/repository"

	"gorm.io/gorm"
)

type GormUserRepository struct {
	*GenericRepository[model.User, repository.CreateUserInput]
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{
		NewGenericRepository(
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

func (r *GormUserRepository) UpdatePassword(id uint, updates repository.UpdateUserPasswordInput) error {
	return r.UpdateByID(id, updates)
}

func (r *GormUserRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	result := r.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *GormUserRepository) UpdateRefreshToken(id uint, updates repository.UpdateUserRefreshTokenInput) error {
	return r.UpdateByID(id, updates)
}
