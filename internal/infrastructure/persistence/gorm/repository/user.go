package repository

import (
	"app/internal/domain/user"

	"gorm.io/gorm"
)

type GormUserRepository struct {
	*Repository[*user.User]
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{
		NewRepository[*user.User](db),
	}
}

func (r *GormUserRepository) FindByEmail(email string) (*user.User, error) {
	var user user.User
	result := r.Repository.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
