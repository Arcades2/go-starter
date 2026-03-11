package auth

import (
	"app/internal/application/auth/contracts"
	"app/internal/application/common"
	"app/internal/domain/user"
)

type AuthService interface {
	common.Panicable
	Login(cmd LoginCommand) (*LoginOutput, error)
	Register(cmd RegisterCommand) (*user.User, error)
}

type authService struct {
	common.BaseService
	userRepo       user.UserRepository
	passwordHasher contracts.PasswordHasher
	tokenGenerator contracts.TokenGenerator
}

func NewAuthService(
	userRepo user.UserRepository,
	passwordHasher contracts.PasswordHasher,
	tokenGenerator contracts.TokenGenerator,
	opts ...common.Option[AuthService],
) AuthService {
	s := &authService{
		userRepo:       userRepo,
		passwordHasher: passwordHasher,
		tokenGenerator: tokenGenerator,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}
