package services

import (
	"app/internal/domain/contracts"
	irepository "app/internal/domain/repository"
	"app/internal/domain/services/authservice"
	"app/internal/domain/services/baseservice"
	"app/internal/domain/services/postservice"
	"app/internal/domain/services/userreaderservice"
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

// REPOSITORIES
func (c *Container) GetUserRepository() irepository.UserRepository {
	return repository.NewGormUserRepository(c.DB)
}

func (c *Container) GetPostRepository() irepository.PostRepository {
	return repository.NewGormPostRepository(c.DB)
}

// INFRASTRUCTURE
func (c *Container) GetPasswordHasher() contracts.PasswordHasher {
	return auth.NewPasswordHasher()
}

func (c *Container) GetTokenGenerator() contracts.TokenGenerator {
	return auth.NewTokenGenerator()
}

// SERVICES
func (c *Container) GetUserService(settings *ServiceSettings) userservice.UserService {
	var opts []baseservice.Option[userservice.UserService]

	if settings != nil && settings.PanicOnError {
		opts = append(opts, baseservice.WithPanicOnError[userservice.UserService])
	}

	return userservice.NewUserService(
		c.GetUserRepository(),
		c.GetPasswordHasher(),
		c.GetUserReaderService(settings),
		opts...,
	)
}

func (c *Container) GetAuthService(settings *ServiceSettings) authservice.AuthService {
	var opts []baseservice.Option[authservice.AuthService]

	if settings != nil && settings.PanicOnError {
		opts = append(opts, baseservice.WithPanicOnError[authservice.AuthService])
	}

	return authservice.NewAuthService(
		c.GetUserRepository(),
		c.GetPasswordHasher(),
		c.GetTokenGenerator(),
		opts...,
	)
}

func (c *Container) GetUserReaderService(settings *ServiceSettings) userreaderservice.UserReader {
	var opts []baseservice.Option[userreaderservice.UserReader]

	if settings != nil && settings.PanicOnError {
		opts = append(opts, baseservice.WithPanicOnError[userreaderservice.UserReader])
	}

	return userreaderservice.NewUserReaderService(
		c.GetUserRepository(),
		opts...,
	)
}

func (c *Container) GetPostService(settings *ServiceSettings) postservice.PostService {
	var opts []baseservice.Option[postservice.PostService]

	if settings != nil && settings.PanicOnError {
		opts = append(opts, baseservice.WithPanicOnError[postservice.PostService])
	}

	return postservice.NewPostService(
		c.GetPostRepository(),
		c.GetUserReaderService(settings),
		opts...,
	)
}
