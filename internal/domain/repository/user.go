package repository

import "app/internal/domain/model"

type UserRepository interface {
	GenericRepository[model.User, CreateUserInput]
	FindByEmail(email string) (*model.User, error)
	UpdateRefreshToken(id uint, refreshToken string) error
	UpdateUserPassword(id uint, updates UpdateUserPasswordInput) error
}

type CreateUserInput struct {
	Firstname      string
	Lastname       string
	Email          string
	HashedPassword string
}

type UpdateUserPasswordInput struct {
	HashedPassword string
}
