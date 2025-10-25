package services

import (
	"app/internal/domain/services/authservice"
	"app/internal/domain/services/userservice"
	"app/internal/infrastructure/auth"
	"app/internal/infrastructure/gorm/repository"

	"gorm.io/gorm"
)

type Container struct {
	DB *gorm.DB
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
func (c *Container) GetUserService() *userservice.UserService {
	return userservice.NewUserService(
		c.GetUserRepository(),
		c.GetPasswordHasher(),
	)
}

func (c *Container) GetAuthService() *authservice.AuthService {
	return authservice.NewAuthService(
		c.GetUserRepository(),
		c.GetPasswordHasher(),
		c.GetTokenGenerator(),
	)
}
