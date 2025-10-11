package userservice

import (
	"app/internal/domain/contracts"
	"app/internal/domain/repository"
)

type UserService struct {
	UserRepository repository.UserRepository
	PasswordHasher contracts.PasswordHasher
}

func NewUserService(
	userRepository repository.UserRepository,
	passwordHasher contracts.PasswordHasher,
) *UserService {
	return &UserService{
		UserRepository: userRepository,
		PasswordHasher: passwordHasher,
	}
}
