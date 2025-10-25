package services

import (
	"app/internal/domain/services/authservice"
	"app/internal/domain/services/baseservice"
	"app/internal/domain/services/userservice"
	"app/internal/infrastructure/auth"
	"app/internal/infrastructure/gorm/repository"

	"gorm.io/gorm"
)

type Container struct {
	DB *gorm.DB
}

type ServiceSettings struct {
	PanicOnError bool
}

func NewContainer(db *gorm.DB) *Container {
	return &Container{
		DB: db,
	}
}

func (c *Container) GetUserRepository() *repository.GormUserRepository {
	return repository.NewGormUserRepository(c.DB)
}

func (c *Container) GetPasswordHasher() *auth.PasswordHasher {
	return auth.NewPasswordHasher()
}

func (c *Container) GetTokenGenerator() *auth.TokenGenerator {
	return auth.NewTokenGenerator()
}

// SERVICES
func (c *Container) GetUserService(settings *ServiceSettings) *userservice.UserService {
	var opts []baseservice.Option[*userservice.UserService]

	if settings != nil && settings.PanicOnError {
		opts = append(opts, baseservice.WithPanicOnError[*userservice.UserService])
	}

	return userservice.NewUserService(
		c.GetUserRepository(),
		c.GetPasswordHasher(),
		opts...,
	)
}

func (c *Container) GetAuthService(settings *ServiceSettings) *authservice.AuthService {
	var opts []baseservice.Option[*authservice.AuthService]

	if settings != nil && settings.PanicOnError {
		opts = append(opts, baseservice.WithPanicOnError[*authservice.AuthService])
	}

	return authservice.NewAuthService(
		c.GetUserRepository(),
		c.GetPasswordHasher(),
		c.GetTokenGenerator(),
		opts...,
	)
}
