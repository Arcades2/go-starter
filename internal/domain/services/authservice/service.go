package authservice

import (
	"app/internal/domain/contracts"
	"app/internal/domain/repository"
)

type AuthService struct {
	UserRepo       repository.UserRepository
	PasswordHasher contracts.PasswordHasher
	TokenGenerator contracts.TokenGenerator
}

func NewAuthService(
	userRepo repository.UserRepository,
	passwordHasher contracts.PasswordHasher,
	tokenGenerator contracts.TokenGenerator,
) *AuthService {
	return &AuthService{
		UserRepo:       userRepo,
		PasswordHasher: passwordHasher,
		TokenGenerator: tokenGenerator,
	}
}
