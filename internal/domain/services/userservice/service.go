package userservice

import (
	"app/internal/domain/contracts"
	"app/internal/domain/repository"
	"app/internal/domain/services/baseservice"
	"app/internal/domain/services/userreaderservice"
)

type UserService interface {
	baseservice.Panicable
	userreaderservice.UserReader
	UpdatePassword(cmd UpdatePasswordCommand) error
}

type userService struct {
	baseservice.BaseService
	userreaderservice.UserReader
	UserRepository repository.UserRepository
	PasswordHasher contracts.PasswordHasher
}

func NewUserService(
	userRepository repository.UserRepository,
	passwordHasher contracts.PasswordHasher,
	reader userreaderservice.UserReader,
	opts ...baseservice.Option[UserService],
) UserService {
	s := &userService{
		UserReader:     reader,
		UserRepository: userRepository,
		PasswordHasher: passwordHasher,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *userService) SetPanicOnError(enable bool) {
	s.BaseService.SetPanicOnError(enable)
}
