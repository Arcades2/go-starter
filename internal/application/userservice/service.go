package userservice

import (
	"app/internal/domain/repository"
)

type UserService struct {
	userRepository repository.UserRepository
	passwordHasher *PasswordHasher
}

func NewUserService(
	userRepository repository.UserRepository,
) *UserService {
	return &UserService{
		userRepository: userRepository,
		passwordHasher: NewPasswordHasher(),
	}
}
