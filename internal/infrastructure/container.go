package infrastructure

import (
	aauth "app/internal/application/auth"
	authcontracts "app/internal/application/auth/contracts"
	apost "app/internal/application/post"
	auser "app/internal/application/user"
	dpost "app/internal/domain/post"
	duser "app/internal/domain/user"
	"app/internal/infrastructure/auth/bcrypt"
	"app/internal/infrastructure/auth/jwt"
	gormrepo "app/internal/infrastructure/persistence/gorm/repository"

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

// USER
func (c *Container) GetUserRepository() duser.UserRepository {
	return gormrepo.NewGormUserRepository(c.DB)
}

func (c *Container) GetUserReaderService() auser.UserReader {
	return auser.NewUserReaderService(
		c.GetUserRepository(),
	)
}

// AUTH
func (c *Container) GetPasswordHasher() authcontracts.PasswordHasher {
	return bcrypt.NewPasswordHasher()
}

func (c *Container) GetTokenGenerator() authcontracts.TokenGenerator {
	return jwt.NewTokenGenerator()
}

func (c *Container) GetAuthService() aauth.AuthService {
	return aauth.NewAuthService(
		c.GetUserRepository(),
		c.GetPasswordHasher(),
		c.GetTokenGenerator(),
	)
}

// POST
func (c *Container) GetPostRepository() dpost.PostRepository {
	return gormrepo.NewGormPostRepository(c.DB)
}

func (c *Container) GetPostReaderService() apost.PostReaderService {
	return apost.NewPostReaderService(
		c.GetPostRepository(),
	)
}

func (c *Container) GetPostService() apost.PostService {
	return apost.NewPostService(
		c.GetPostRepository(),
		c.GetPostReaderService(),
		c.GetUserReaderService(),
	)
}
