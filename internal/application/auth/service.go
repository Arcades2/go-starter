package auth

import (
	"app/internal/application/auth/contracts"
	"app/internal/domain/user"
)

type AuthService interface {
	Login(cmd LoginCommand) (*LoginOutput, error)
	Register(cmd RegisterCommand) (*user.User, error)
}

type authService struct {
	userRepo       user.UserRepository
	passwordHasher contracts.PasswordHasher
	tokenGenerator contracts.TokenGenerator
}

func NewAuthService(
	userRepo user.UserRepository,
	passwordHasher contracts.PasswordHasher,
	tokenGenerator contracts.TokenGenerator,
) AuthService {
	return &authService{
		userRepo:       userRepo,
		passwordHasher: passwordHasher,
		tokenGenerator: tokenGenerator,
	}
}
