package authservice

import (
	"app/internal/domain/contracts"
	"app/internal/domain/repository"
	"app/internal/domain/services/baseservice"
)

type AuthService struct {
	baseservice.BaseService
	UserRepo       repository.UserRepository
	PasswordHasher contracts.PasswordHasher
	TokenGenerator contracts.TokenGenerator
}

func NewAuthService(
	userRepo repository.UserRepository,
	passwordHasher contracts.PasswordHasher,
	tokenGenerator contracts.TokenGenerator,
	opts ...baseservice.Option[*AuthService],
) *AuthService {
	s := &AuthService{
		UserRepo:       userRepo,
		PasswordHasher: passwordHasher,
		TokenGenerator: tokenGenerator,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}
