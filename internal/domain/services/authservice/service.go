package authservice

import (
	"app/internal/domain/contracts"
	"app/internal/domain/model"
	"app/internal/domain/repository"
	"app/internal/domain/services/baseservice"
)

type AuthService interface {
	baseservice.Panicable
	Login(cmd LoginCommand) (*LoginOutput, error)
	Register(cmd RegisterCommand) (*model.User, error)
}

type authService struct {
	baseservice.BaseService
	UserRepo       repository.UserRepository
	PasswordHasher contracts.PasswordHasher
	TokenGenerator contracts.TokenGenerator
}

func NewAuthService(
	userRepo repository.UserRepository,
	passwordHasher contracts.PasswordHasher,
	tokenGenerator contracts.TokenGenerator,
	opts ...baseservice.Option[AuthService],
) AuthService {
	s := &authService{
		UserRepo:       userRepo,
		PasswordHasher: passwordHasher,
		TokenGenerator: tokenGenerator,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}
