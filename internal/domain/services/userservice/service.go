package userservice

import (
	"app/internal/domain/contracts"
	"app/internal/domain/repository"
	"app/internal/domain/services/baseservice"
)

type UserService struct {
	baseservice.BaseService
	UserRepository repository.UserRepository
	PasswordHasher contracts.PasswordHasher
}

func NewUserService(
	userRepository repository.UserRepository,
	passwordHasher contracts.PasswordHasher,
	opts ...baseservice.Option[*UserService],
) *UserService {
	s := &UserService{
		UserRepository: userRepository,
		PasswordHasher: passwordHasher,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}
