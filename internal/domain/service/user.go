package service

import "app/internal/domain/model"

type UserServiceInterface interface {
	CreateUser(data CreateUserCommand) (*model.User, error)
	GetUserByID(ID uint) (*model.User, error)
	UpdateUserPassword(ID uint, newPassword string) error
	UpdateUserInfo(ID uint, firstname, lastname string) error
}

type CreateUserCommand struct {
	Firstname string `validate:"required,min=1,max=255"`
	Lastname  string `validate:"required,min=1,max=255"`
	Email     string `validate:"required,email,max=255"`
	Password  string `validate:"required,min=8,max=100"`
}
