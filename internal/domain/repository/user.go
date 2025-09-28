package repository

import "app/internal/domain/model"

type UserRepository interface {
	GenericRepository[model.User, CreateUserInput]
	UpdateUserInfo(id uint, updates UpdateUserInfoInput) error
	UpdateUserPassword(id uint, updates UpdateUserPasswordInput) error
}

type CreateUserInput struct {
	Firstname      string
	Lastname       string
	Email          string
	HashedPassword string
}

type UpdateUserInfoInput struct {
	Firstname string
	Lastname  string
}

type UpdateUserPasswordInput struct {
	HashedPassword string
}
